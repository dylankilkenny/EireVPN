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

// Tokens creates a jwt token from the user ID. This token will
// expire in 1 hour
func Tokens(usersession models.UserAppSession) (string, string, string, error) {
	conf := config.GetConfig()

	// Set claims
	csrfExpiry := jwt_lib.StandardClaims{
		ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
	}
	csrfTokenClaims := JWTClaims{
		StandardClaims: csrfExpiry,
	}
	// Create csrf token
	csrfToken := jwt_lib.NewWithClaims(jwt_lib.GetSigningMethod("HS256"), csrfTokenClaims)
	csrfTokenString, err := csrfToken.SignedString([]byte(conf.App.JWTSecret))

	authExpiry := jwt_lib.StandardClaims{
		ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
	}
	authTokenClaims := JWTClaims{
		UserID:         usersession.UserID,
		CSRF:           csrfTokenString,
		StandardClaims: authExpiry,
	}
	refreshExpiry := jwt_lib.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	}
	refreshTokenClaims := JWTClaims{
		UserID:            usersession.UserID,
		SessionIdentifier: usersession.Identifier,
		StandardClaims:    refreshExpiry,
	}
	// Create auth token
	authToken := jwt_lib.NewWithClaims(jwt_lib.GetSigningMethod("HS256"), authTokenClaims)
	// Create refresh token
	refreshToken := jwt_lib.NewWithClaims(jwt_lib.GetSigningMethod("HS256"), refreshTokenClaims)

	// Sign and get the complete encoded token as a string
	authTokenString, err := authToken.SignedString([]byte(conf.App.JWTSecret))
	refreshTokenString, err := refreshToken.SignedString([]byte(conf.App.JWTSecret))
	if err != nil {
		return "", "", "", err
	}
	return authTokenString, refreshTokenString, csrfTokenString, nil
}

// PasswordResetToken creates a one time us jwt token from the users old password
func PasswordResetToken(id string) (string, error) {
	conf := config.GetConfig()
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
	conf := config.GetConfig()
	token, err := jwt_lib.ParseWithClaims(refreshToken, &JWTClaims{}, func(token *jwt_lib.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt_lib.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(conf.App.JWTSecret), nil
	})

	if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}

// ValidateString Validate token string
func ValidateString(token string) (bool, error) {
	conf := config.GetConfig()
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
