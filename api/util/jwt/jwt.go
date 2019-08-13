package jwt

import (
	"eirevpn/api/models"
	"eirevpn/api/util/random"
	"fmt"
	"time"

	jwt_lib "github.com/dgrijalva/jwt-go"
)

// TODO: add this to env variables
const secretkey = "verysecretkey1995"

type JWTClaims struct {
	UserID            uint   `json:"user_id"`
	CSRF              string `json:"csrf"`
	SessionIdentifier string `json:"Identifier"`
	jwt_lib.StandardClaims
}

// Tokens creates a jwt token from the user ID. This token will
// expire in 1 hour
func Tokens(usersession models.UserSession) (string, string, string, error) {
	// Set claims
	csrfToken, _ := random.GenerateRandomString(64)
	authExpiry := jwt_lib.StandardClaims{
		ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
	}
	authTokenClaims := JWTClaims{
		UserID:         usersession.UserID,
		CSRF:           csrfToken,
		StandardClaims: authExpiry,
	}
	refreshExpiry := jwt_lib.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
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
	authTokenString, err := authToken.SignedString([]byte(secretkey))
	refreshTokenString, err := refreshToken.SignedString([]byte(secretkey))
	if err != nil {
		return "", "", "", err
	}
	return authTokenString, refreshTokenString, csrfToken, nil
}

// PasswordResetToken creates a one time us jwt token from the users old password
func PasswordResetToken(id string) (string, error) {
	// Create the token
	token := jwt_lib.New(jwt_lib.GetSigningMethod("HS256"))
	// Set some claims
	token.Claims = jwt_lib.MapClaims{
		"Id":  id,
		"exp": time.Now().Add(time.Hour * 1).Unix(),
	}
	// Sign and get the complete encoded token as a string
	tokenString, err := token.SignedString([]byte(secretkey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// ValidateAuthToken todo
func ValidateToken(refreshToken string) (*JWTClaims, error) {

	token, err := jwt_lib.ParseWithClaims(refreshToken, &JWTClaims{}, func(token *jwt_lib.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt_lib.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretkey), nil
	})

	if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}

// ValidateString Validate token string
func ValidateString(token string) (bool, error) {
	_, err := jwt_lib.Parse(token, func(token *jwt_lib.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt_lib.SigningMethodHMAC); !ok {
			return false, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretkey), nil
	})
	if err != nil {
		return false, err
	}
	return true, nil
}
