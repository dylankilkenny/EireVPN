package settings

import (
	"eirevpn/api/config"
	"eirevpn/api/errors"
	"eirevpn/api/logger"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SettingsFields struct {
	EnableCSRF          string `json:"enableCsrf" binding:"required"`
	EnableSubscriptions string `json:"enableSubscriptions" binding:"required"`
	EnableAuth          string `json:"enableAuth" binding:"required"`
	IntegrationActive   string `json:"enableStripe" binding:"required"`
}

// UpdateSettings updates the config.yaml
func UpdateSettings(c *gin.Context) {

	settingsUpdates := SettingsFields{}
	if err := c.BindJSON(&settingsUpdates); err != nil {
		logger.Log(logger.Fields{
			Loc:  "/settings/update - UpdateSettings()",
			Code: errors.SettingsUpdateFailed.Code,
			Err:  err.Error(),
		})
		c.AbortWithStatusJSON(errors.SettingsUpdateFailed.Status, errors.SettingsUpdateFailed)
		return
	}

	conf := config.GetConfig()
	conf.App.EnableCSRF = settingsUpdates.EnableCSRF == "true"
	conf.App.EnableSubscriptions = settingsUpdates.EnableSubscriptions == "true"
	conf.App.EnableAuth = settingsUpdates.EnableAuth == "true"
	conf.Stripe.IntegrationActive = settingsUpdates.IntegrationActive == "true"

	if err := conf.SaveConfig(); err != nil {
		logger.Log(logger.Fields{
			Loc:   "/settings/update - UpdateSettings()",
			Code:  errors.SettingsUpdateFailed.Code,
			Extra: map[string]interface{}{"PlanID": c.Param("id")},
			Err:   err.Error(),
		})
		c.AbortWithStatusJSON(errors.SettingsUpdateFailed.Status, errors.SettingsUpdateFailed)
		return

	}
	c.JSON(http.StatusOK, gin.H{
		"status": 200,
	})
}

// Settings fetches the settings
func Settings(c *gin.Context) {
	conf := config.GetConfig()

	s := SettingsFields{
		EnableCSRF:          strconv.FormatBool(conf.App.EnableCSRF),
		EnableSubscriptions: strconv.FormatBool(conf.App.EnableSubscriptions),
		EnableAuth:          strconv.FormatBool(conf.App.EnableAuth),
		IntegrationActive:   strconv.FormatBool(conf.Stripe.IntegrationActive),
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"data": gin.H{
			"settings": s,
		},
	})

}
