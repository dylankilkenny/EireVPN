package jwt

import (
	"eirevpn/api/config"
	"eirevpn/api/models"
	"fmt"
	"time"

	jwt_lib "github.com/dgrijalva/jwt-go"
)

type JWTClaims struct {
	UserID            uint   `json:"user_id"`
	CSRF              string `json:"csrf"`
	SessionIdentifier string `json:"Identifier"`
	jwt_lib.StandardClaims
}

// Tokens creates a jwt token from the user ID
func Tokens(usersession models.UserAppSession) (string, string, string, error) {
	conf := config.Load()
	authExpiry := time.Hour * time.Duration(conf.App.AuthTokenExpiry)
	refreshExpiry := time.Hour * time.Duration(conf.App.AuthTokenExpiry)

	// Set claims
	csrfTokenClaims := JWTClaims{
		StandardClaims: jwt_lib.StandardClaims{
			ExpiresAt: time.Now().Add(authExpiry).Unix(),
		},
	}
	// Create csrf token
	csrfToken := jwt_lib.NewWithClaims(jwt_lib.GetSigningMethod("HS256"), csrfTokenClaims)
	csrfTokenString, err := csrfToken.SignedString([]byte(conf.App.JWTSecret))

	authTokenClaims := JWTClaims{
		UserID:            usersession.UserID,
		CSRF:              csrfTokenString,
		SessionIdentifier: usersession.Identifier,
		StandardClaims: jwt_lib.StandardClaims{
			ExpiresAt: time.Now().Add(authExpiry).Unix(),
		},
	}
	// Create refresh token
	authToken := jwt_lib.NewWithClaims(jwt_lib.GetSigningMethod("HS256"), authTokenClaims)
	// Sign and get the complete encoded token as a string
	authTokenString, err := authToken.SignedString([]byte(conf.App.JWTSecret))

	refreshTokenClaims := JWTClaims{
		UserID: usersession.UserID,
		StandardClaims: jwt_lib.StandardClaims{
			ExpiresAt: time.Now().Add(refreshExpiry).Unix(),
		},
	}
	// Create refresh token
	refreshToken := jwt_lib.NewWithClaims(jwt_lib.GetSigningMethod("HS256"), refreshTokenClaims)
	// Sign and get the complete encoded token as a string
	refreshTokenString, err := refreshToken.SignedString([]byte(conf.App.JWTSecret))

	if err != nil {
		return "", "", "", err
	}

	return authTokenString, refreshTokenString, csrfTokenString, nil
}

// PasswordResetToken creates a one time us jwt token from the users old password
func PasswordResetToken(id string) (string, error) {
	conf := config.Load()
	// Create the token
	token := jwt_lib.New(jwt_lib.GetSigningMethod("HS256"))
	// Set some claims
	token.Claims = jwt_lib.MapClaims{
		"Id":  id,
		"exp": time.Now().Add(time.Hour * 1).Unix(),
	}
	// Sign and get the complete encoded token as a string
	tokenString, err := token.SignedString([]byte(conf.App.JWTSecret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// ValidateAuthToken todo
func ValidateToken(refreshToken string) (*JWTClaims, error) {
	conf := config.Load()
	token, err := jwt_lib.ParseWithClaims(refreshToken, &JWTClaims{}, func(token *jwt_lib.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt_lib.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(conf.App.JWTSecret), nil
	})
	claims, ok := token.Claims.(*JWTClaims)
	if ok && token.Valid {
		return claims, nil
	}
	return claims, err
}

// ValidateString Validate token string
func ValidateString(token string) (bool, error) {
	conf := config.Load()
	_, err := jwt_lib.Parse(token, func(token *jwt_lib.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt_lib.SigningMethodHMAC); !ok {
			return false, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(conf.App.JWTSecret), nil
	})
	if err != nil {
		return false, err
	}
	return true, nil
}
