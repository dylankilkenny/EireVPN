package message

import (
	"eirevpn/api/errors"
	"eirevpn/api/integrations/sendgrid"
	"eirevpn/api/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MessageFields struct {
	Email   string `json:"email" binding:"required"`
	Subject string `json:"subject" binding:"required"`
	Message string `json:"message" binding:"required"`
}

// Message rel
func Message(c *gin.Context) {

	mf := MessageFields{}
	if err := c.BindJSON(&mf); err != nil {
		logger.Log(logger.Fields{
			Loc:  "/message - Message()",
			Code: errors.MsgBindingFailed.Code,
			Err:  err.Error(),
		})
		c.AbortWithStatusJSON(errors.InternalServerError.Status, errors.InternalServerError)
		return
	}

	if err := sendgrid.Send().SupportRequest(mf.Email, mf.Subject, mf.Message); err != nil {
		logger.Log(logger.Fields{
			Loc:   "/message - Message()",
			Code:  errors.InternalServerError.Code,
			Extra: map[string]interface{}{"Email": mf.Email, "Detail": "Error sending support email"},
			Err:   err.Error(),
		})
		c.AbortWithStatusJSON(errors.InternalServerError.Status, errors.InternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 200,
	})
}
