package user

import (
	"eirevpn/api/logger"
	"net/http"
	"time"

	"eirevpn/api/db"
	"eirevpn/api/errors"
	"eirevpn/api/models"
	"eirevpn/api/util/jwt"

	"github.com/gin-gonic/gin"
	logrus "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

var t = time.Now()

type token struct {
	Token string `json:"token" binding:"required"`
}

// LoginUser verifies a users details are correct, returning a jwt token to the user
func LoginUser(c *gin.Context) {
	db := db.GetDB()

	var userLogin models.User
	var userDb models.User

	if err := c.BindJSON(&userLogin); err != nil {
		logger.Log(logrus.Fields{
			"Loc:":   "/login - LoginUser()",
			"Code:":  errors.EmailOrPassword.Code,
			"Email:": userLogin.Email,
		}, err.Error())
		c.AbortWithStatusJSON(errors.EmailOrPassword.Status, errors.EmailOrPassword)
		return
	}

	if err := db.Where("email = ?", userLogin.Email).First(&userDb).Error; err != nil {
		logger.Log(logrus.Fields{
			"Loc:":   "/login - LoginUser()",
			"Code:":  errors.EmailNotFound.Code,
			"Email:": userLogin.Email,
		}, err.Error())
		c.AbortWithStatusJSON(errors.EmailNotFound.Status, errors.EmailNotFound)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userDb.Password), []byte(userLogin.Password)); err != nil {
		logger.Log(logrus.Fields{
			"Loc:":   "/login - LoginUser()",
			"Code:":  errors.WrongPassword.Code,
			"Email:": userLogin.Email,
		}, err.Error())
		c.AbortWithStatusJSON(errors.WrongPassword.Status, errors.WrongPassword)
		return
	}

	var usersession models.UserSession
	usersession.UserID = userDb.ID
	if err := db.Create(&usersession).Error; err != nil {
		logger.Log(logrus.Fields{
			"Loc:":    "/login - LoginUser()",
			"Code:":   errors.InternalServerError.Code,
			"UserID:": userDb.ID,
		}, err.Error())
		c.AbortWithStatusJSON(errors.InternalServerError.Status, errors.InternalServerError)
		return
	}

	authToken, refreshToken, csrfToken, err := jwt.Token(usersession)
	if err != nil {
		logger.Log(logrus.Fields{
			"Loc:":    "/login - LoginUser()",
			"Code:":   errors.InternalServerError.Code,
			"UserID:": userDb.ID,
		}, err.Error())
		c.AbortWithStatusJSON(errors.InternalServerError.Status, errors.InternalServerError)
		return
	}

	// TODO: Change the domain name and add correct maxAge time
	authCookieMaxAge := 15 * 60 // 15 minutes in seconds
	c.SetCookie("authToken", authToken, authCookieMaxAge, "/", "localhost", true, false)

	// TODO: Change the domain name and add correct maxAge time
	refreshCookieMaxAge := 72 * 60 * 60 // 72 hours in seconds
	c.SetCookie("refreshToken", refreshToken, refreshCookieMaxAge, "/", "localhost", true, false)

	c.Header("X-CSRF-Token", csrfToken)
	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"errors": make([]string, 0),
		"data":   gin.H{"firstname": userDb.FirstName},
	})
}

// SignUpUser registers a new user
func SignUpUser(c *gin.Context) {
	db := db.GetDB()
	var user models.User

	if err := c.BindJSON(&user); err != nil {
		logger.Log(logrus.Fields{
			"Loc:":   "/signup - SignUpUser()",
			"Code:":  errors.InvalidForm.Code,
			"Email:": user.Email,
		}, err.Error())
		c.AbortWithStatusJSON(errors.InvalidForm.Status, errors.InvalidForm)
		return
	}

	if err := db.Where("email = ?", user.Email).First(&user).Error; err == nil {
		logger.Log(logrus.Fields{
			"Loc:":   "/signup - SignUpUser()",
			"Code:":  errors.EmailTaken.Code,
			"Email:": user.Email,
		}, err.Error())
		c.JSON(errors.EmailTaken.Status, errors.EmailTaken)
		return
	}

	if err := db.Create(&user).Error; err != nil {
		logger.Log(logrus.Fields{
			"Loc:":   "/signup - SignUpUser()",
			"Code:":  errors.InternalServerError.Code,
			"Email:": user.Email,
		}, err.Error())
		c.AbortWithStatusJSON(errors.InternalServerError.Status, errors.InternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"data":   make([]string, 0),
	})
}

// ChangePasswordRequest sends the user a link to change their password
func ChangePasswordRequest(c *gin.Context) {
	db := db.GetDB()
	email := c.PostForm("email")
	var user models.User

	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
		c.AbortWithStatusJSON(errors.EmailNotFound.Status, errors.EmailNotFound)
		return
	}

	// Send email here

	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"errors": make([]string, 0),
		"data":   make([]string, 0),
	})
}

// ChangePassword will authenticate the users token and change their password
func ChangePassword(c *gin.Context) {
	db := db.GetDB()
	email := c.PostForm("email")
	var user models.User

	// _, err := jwt.Validate(c)
	// if err != nil {
	// 	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
	// 		"status": 401,
	// 		"errors": gin.H{
	// 			"title":  "Invalid Token",
	// 			"detail": "Token provided in auth header is not valid",
	// 		},
	// 		"data": make([]string, 0),
	// 	})
	// 	return
	// }

	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
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

	// Send email here

	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"errors": make([]string, 0),
		"data":   make([]string, 0),
	})
}
