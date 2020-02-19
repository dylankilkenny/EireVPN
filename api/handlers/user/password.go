package user

import (
	"eirevpn/api/errors"
	"eirevpn/api/integrations/sendgrid"
	"eirevpn/api/logger"
	"eirevpn/api/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// ForgotPasswordToken will email the user a token which will
// allow them to change their password
func ForgotPasswordToken(c *gin.Context) {
	type User struct {
		Email string `json:"email" binding:"required"`
	}

	u := User{}
	if err := c.BindJSON(&u); err != nil {
		logger.Log(logger.Fields{
			Loc:  "/forgot_pass - ForgotPasswordToken()",
			Code: errors.BindingFailed.Code,
			Err:  err.Error(),
		})
		c.AbortWithStatusJSON(errors.InternalServerError.Status, errors.InternalServerError)
		return
	}

	var user models.User
	user.Email = u.Email
	if err := user.Find(); err != nil {
		logger.Log(logger.Fields{
			Loc:   "/forgot_pass - ForgotPasswordToken()",
			Code:  errors.UserNotFound.Code,
			Extra: map[string]interface{}{"Email": u.Email},
			Err:   errors.UserNotFound.Detail,
		})
		// return with ok status as we dont want to leak which emails
		// have an account
		c.JSON(http.StatusOK, gin.H{"status": 200})
		return
	}

	var fp models.ForgotPassword
	fp.UserID = user.ID
	if err := fp.Create(); err != nil {
		logger.Log(logger.Fields{
			Loc:   "/forgot_pass - ForgotPasswordToken()",
			Code:  errors.InternalServerError.Code,
			Extra: map[string]interface{}{"UserID": user.ID, "Detail": "Error creating forgot password object"},
			Err:   err.Error(),
		})
		c.AbortWithStatusJSON(errors.InternalServerError.Status, errors.InternalServerError)
		return
	}

	if err := sendgrid.Send().ForgotPassword(u.Email, fp.Token); err != nil {
		logger.Log(logger.Fields{
			Loc:   "/forgot_pass - ForgotPasswordToken()",
			Code:  errors.InternalServerError.Code,
			Extra: map[string]interface{}{"Email": u.Email, "Detail": "Error sending support email"},
			Err:   err.Error(),
		})
		c.AbortWithStatusJSON(errors.InternalServerError.Status, errors.InternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 200,
	})
}

// UpdatePassword will verify the token is invalid and update the password
// for the user associated with the token
func UpdatePassword(c *gin.Context) {

	type User struct {
		Password string `json:"password" binding:"required"`
	}

	token := c.Param("token")

	var fp models.ForgotPassword
	fp.Token = token
	if err := fp.Find(); err != nil {
		logger.Log(logger.Fields{
			Loc:   "/forgot_pass - UpdatePassword()",
			Code:  errors.TokenNotFound.Code,
			Extra: map[string]interface{}{"Token": c.Param("token")},
			Err:   err.Error(),
		})
		c.AbortWithStatusJSON(errors.TokenNotFound.Status, errors.TokenNotFound)
		return
	}

	u := User{}
	if err := c.BindJSON(&u); err != nil {
		logger.Log(logger.Fields{
			Loc:  "/forgot_pass - UpdatePassword()",
			Code: errors.BindingFailed.Code,
			Err:  err.Error(),
		})
		c.AbortWithStatusJSON(errors.InternalServerError.Status, errors.InternalServerError)
		return
	}

	var user models.User
	user.ID = fp.UserID
	if err := user.Find(); err != nil {
		logger.Log(logger.Fields{
			Loc:   "/forgot_pass - UpdatePassword()",
			Code:  errors.UserNotFound.Code,
			Extra: map[string]interface{}{"UserID": fp.UserID},
			Err:   errors.UserNotFound.Detail,
		})
		c.JSON(errors.UserNotFound.Status, errors.UserNotFound)
		return
	}

	pw, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		logger.Log(logger.Fields{
			Loc:   "/forgot_pass - UpdatePassword()",
			Code:  errors.InternalServerError.Code,
			Extra: map[string]interface{}{"UserID": user.ID},
			Err:   err.Error(),
		})
		c.AbortWithStatusJSON(errors.InternalServerError.Status, errors.InternalServerError)
		return
	}

	user.Password = string(pw)
	if err := user.Save(); err != nil {
		logger.Log(logger.Fields{
			Loc:   "/forgot_pass - UpdatePassword()",
			Code:  errors.InternalServerError.Code,
			Extra: map[string]interface{}{"UserID": user.ID},
			Err:   err.Error(),
		})
		c.AbortWithStatusJSON(errors.InternalServerError.Status, errors.InternalServerError)
		return
	}

	if err := fp.Delete(); err != nil {
		logger.Log(logger.Fields{
			Loc:   "/forgot_pass - UpdatePassword()",
			Code:  errors.UserNotFound.Code,
			Extra: map[string]interface{}{"UserID": user.ID},
			Err:   err.Error(),
		})
		c.AbortWithStatusJSON(errors.UserNotFound.Status, errors.UserNotFound)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 200,
	})

}

// ChangePassword will authenticate the users token and change their password
func ChangePassword(c *gin.Context) {

	cookieUserID, exists := c.Get("UserID")
	if !exists {
		logger.Log(logger.Fields{
			Loc: "/user/private/changepassword - ChangePassword()",
			Extra: map[string]interface{}{
				"UserID": cookieUserID,
				"Detail": "User ID does not exist in the context",
			},
		})
		c.AbortWithStatusJSON(errors.InternalServerError.Status, errors.InternalServerError)
		return
	}

	var user models.User
	user.ID = cookieUserID.(uint)
	if err := user.Find(); err != nil {
		logger.Log(logger.Fields{
			Loc:   "/user/private/changepassword - ChangePassword()",
			Code:  errors.UserNotFound.Code,
			Extra: map[string]interface{}{"UserID": c.Param("id")},
			Err:   err.Error(),
		})
		c.AbortWithStatusJSON(errors.UserNotFound.Status, errors.UserNotFound)
		return
	}

	type ChangePassword struct {
		CurrentPassword string `json:"current_password" binding:"required"`
		NewPassword     string `json:"new_password" binding:"required"`
	}

	changePassword := ChangePassword{}

	if err := c.BindJSON(&changePassword); err != nil {
		logger.Log(logger.Fields{
			Loc:   "/user/private/changepassword - ChangePassword()",
			Code:  errors.InvalidForm.Code,
			Extra: map[string]interface{}{"UserID": cookieUserID},
			Err:   err.Error(),
		})
		c.AbortWithStatusJSON(errors.InvalidForm.Status, errors.InvalidForm)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(changePassword.CurrentPassword)); err != nil {
		logger.Log(logger.Fields{
			Loc:   "/user/private/changepassword - ChangePassword()",
			Code:  errors.WrongPassword.Code,
			Extra: map[string]interface{}{"UserID": cookieUserID},
			Err:   err.Error(),
		})
		c.AbortWithStatusJSON(errors.WrongPassword.Status, errors.WrongPassword)
		return
	}

	pw, err := bcrypt.GenerateFromPassword([]byte(changePassword.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		logger.Log(logger.Fields{
			Loc:   "/user/private/changepassword - ChangePassword()",
			Code:  errors.InternalServerError.Code,
			Extra: map[string]interface{}{"UserID": cookieUserID},
			Err:   err.Error(),
		})
		c.AbortWithStatusJSON(errors.InternalServerError.Status, errors.InternalServerError)
		return
	}

	user.Password = string(pw)
	if err := user.Save(); err != nil {
		logger.Log(logger.Fields{
			Loc:   "/user/private/changepassword - ChangePassword()",
			Code:  errors.InternalServerError.Code,
			Extra: map[string]interface{}{"UserID": user.ID},
			Err:   err.Error(),
		})
		c.AbortWithStatusJSON(errors.InternalServerError.Status, errors.InternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"errors": make([]string, 0),
		"data":   make([]string, 0),
	})
}
