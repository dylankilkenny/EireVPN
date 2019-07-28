package jwt

import (
	"fmt"
	"strings"

	jwt_lib "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// TODO: add this to env variables
const secretkey = "verysecretkey1995"

func Token(id string) (string, error) {
	// Create the token
	token := jwt_lib.New(jwt_lib.GetSigningMethod("HS256"))
	// Set some claims
	token.Claims = jwt_lib.MapClaims{
		"Id": id,
	}
	// Sign and get the complete encoded token as a string
	tokenString, err := token.SignedString([]byte(secretkey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// taken from https://godoc.org/github.com/dgrijalva/jwt-go#example-Parse--Hmac
func Validate(c *gin.Context) (string, error) {
	auth := c.Request.Header.Get("Authorization")
	if auth == "" {
		return "", fmt.Errorf("No token")
	}
	tokenString := strings.Split(auth, " ")[1]
	token, err := jwt_lib.Parse(tokenString, func(token *jwt_lib.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt_lib.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretkey), nil
	})

	if token == nil {
		return "", fmt.Errorf("No token")
	}

	if claims, ok := token.Claims.(jwt_lib.MapClaims); ok && token.Valid {
		return claims["Id"].(string), nil
	}
	return "", err
}

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
