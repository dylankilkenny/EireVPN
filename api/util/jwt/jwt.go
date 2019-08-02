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

// Token creates a jwt token from the user ID. This token will
// expire in 1 hour
func Token(usersession models.UserSession) (string, string, string, error) {
	// Create auth token
	authToken := jwt_lib.New(jwt_lib.GetSigningMethod("HS256"))
	// Create refresh token
	refreshToken := jwt_lib.New(jwt_lib.GetSigningMethod("HS256"))
	// Set some claims
	csrfToken, _ := random.GenerateRandomString(64)
	authToken.Claims = jwt_lib.MapClaims{
		"Id":   usersession.UserID,
		"exp":  time.Now().Add(time.Minute * 15).Unix(),
		"csrf": csrfToken,
	}
	refreshToken.Claims = jwt_lib.MapClaims{
		"Id":         usersession.UserID,
		"Identifier": usersession.Identifier,
		"exp":        time.Now().Add(time.Hour * 72).Unix(),
	}
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
func ValidateAuthToken(authToken string) (jwt_lib.MapClaims, error) {

	token, err := jwt_lib.Parse(authToken, func(token *jwt_lib.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt_lib.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretkey), nil
	})

	if claims, ok := token.Claims.(jwt_lib.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}

// ValidateAuthToken todo
func ValidateRefreshToken(refreshToken string) (jwt_lib.MapClaims, error) {

	token, err := jwt_lib.Parse(refreshToken, func(token *jwt_lib.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt_lib.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretkey), nil
	})

	if claims, ok := token.Claims.(jwt_lib.MapClaims); ok && token.Valid {
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
