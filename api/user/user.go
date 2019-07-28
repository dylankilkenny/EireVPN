package user

import (
	"fmt"
	"net/http"

	"eirevpn/api/db"
	"eirevpn/api/models"
	"eirevpn/api/util/jwt"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type token struct {
	Token string `json:"token" binding:"required"`
}

// LoginUser verifies a users details are correct, returning a jwt token to the user
func LoginUser(c *gin.Context) {
	db := db.GetDB()

	var userLogin models.User
	var userDb models.User

	if err := c.BindJSON(&userLogin); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"errors": gin.H{
				"title":  "Email or password missing",
				"detail": "Email or password missing",
			},
			"data": make([]string, 0),
		})
		return
	}

	if err := db.Where("email = ?", userLogin.Email).First(&userDb).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"status": 404,
			"errors": gin.H{
				"title":  "Email Not Found",
				"detail": "No matching email address was found",
			},
			"data": make([]string, 0),
		})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userDb.Password), []byte(userLogin.Password)); err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"status": 401,
			"errors": gin.H{
				"title":  "Wrong Password",
				"detail": "password is incorrect",
			},
			"data": make([]string, 0),
		})
		return
	}

	token, err := jwt.Token(string(userDb.ID))
	if err != nil {
		fmt.Printf("Error creating token for user %v", userDb.ID)
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"errors": make([]string, 0),
		"data": gin.H{
			"token":     token,
			"firstname": userDb.FirstName,
		},
	})
}

// SignUpUser registers a new user
func SignUpUser(c *gin.Context) {
	db := db.GetDB()
	var user models.User

	if err := c.BindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"errors": gin.H{
				"title":  "Invalid Form",
				"detail": "The submitted form is not valid",
			},
			"data": make([]string, 0),
		})
		return
	}

	if err := db.Where("email = ?", user.Email).First(&user).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"errors": gin.H{
				"title":  "Email Taken",
				"detail": "email already exists",
			},
			"data": make([]string, 0),
		})
		return
	}

	db.Create(&user)
	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"errors": make([]string, 0),
		"data":   make([]string, 0),
	})
}
