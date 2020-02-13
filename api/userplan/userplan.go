package userplan

import (
	"eirevpn/api/config"
	"eirevpn/api/errors"
	"eirevpn/api/logger"
	"eirevpn/api/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// UserPlan fetches a plan by ID
func UserPlan(c *gin.Context) {
	conf := config.Load()

	userID, _ := strconv.ParseUint(c.Param("userid"), 10, 64)
	var userplan models.UserPlan
	userplan.UserID = uint(userID)
	if err := userplan.Find(); err != nil {
		logger.Log(logger.Fields{
			Loc:   "/userplans/:id - UserPlan()",
			Code:  errors.UserPlanNotFound.Code,
			Extra: map[string]interface{}{"UserPlanID": c.Param("id")},
			Err:   err.Error(),
		})
		c.AbortWithStatusJSON(errors.UserPlanNotFound.Status, errors.UserPlanNotFound)
		return
	}

	cookieUserID, exists := c.Get("UserID")
	if !exists {
		logger.Log(logger.Fields{
			Loc: "/userplans/:id - UserPlan()",
			Extra: map[string]interface{}{
				"UserID": userID,
				"Detail": "User ID does not exist in the context",
			},
		})
		c.AbortWithStatusJSON(errors.InternalServerError.Status, errors.InternalServerError)
		return
	}
	if cookieUserID.(uint) != userplan.UserID {
		logger.Log(logger.Fields{
			Loc:  "/userplans/:id - UserPlan()",
			Code: errors.ProtectedRouted.Code,
			Extra: map[string]interface{}{
				"CookieUserID": cookieUserID,
				"QueryUserID":  userID,
			},
			Err: "User does not have permission to access route",
		})

		c.SetCookie(conf.App.AuthCookieName, "", -1, "/", conf.App.Domain, false, true)
		c.SetCookie("uid", "", -1, "/", conf.App.Domain, false, false)
		c.AbortWithStatusJSON(errors.ProtectedRouted.Status, errors.ProtectedRouted)
		return
	}

	type UserPlanCustom struct {
		models.UserPlan
		PlanType models.PlanType `json:"plan_type"`
		PlanName string          `json:"plan_name"`
	}

	var plan models.Plan
	plan.ID = userplan.PlanID
	if err := plan.Find(); err != nil {
		logger.Log(logger.Fields{
			Loc:   "/userplans/:id - UserPlan()",
			Code:  errors.PlanNotFound.Code,
			Extra: map[string]interface{}{"PlanID": userplan.PlanID},
			Err:   err.Error(),
		})
		c.AbortWithStatusJSON(errors.PlanNotFound.Status, errors.PlanNotFound)
		return
	}

	respUserPlan := UserPlanCustom{
		UserPlan: userplan,
		PlanType: plan.PlanType,
		PlanName: plan.Name,
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"errors": make([]string, 0),
		"data": gin.H{
			"userplan": respUserPlan,
		},
	})

}

// CreateUserPlan creates a new plan
func CreateUserPlan(c *gin.Context) {
	var userplan models.UserPlan
	type UserPlanCreate struct {
		UserID     uint   `json:"user_id" binding:"required"`
		PlanID     uint   `json:"plan_id" binding:"required"`
		Active     string `json:"active" binding:"required"`
		StartDate  string `json:"start_date" binding:"required"`
		ExpiryDate string `json:"expiry_date" binding:"required"`
	}
	userPlanCreate := UserPlanCreate{}
	if err := c.BindJSON(&userPlanCreate); err != nil {
		logger.Log(logger.Fields{
			Loc:  "/userplans/create - CreateUserPlan()",
			Code: errors.InvalidForm.Code,
			Err:  err.Error(),
		})
		c.AbortWithStatusJSON(errors.InvalidForm.Status, errors.InvalidForm)
		return
	}

	userplan.UserID = userPlanCreate.UserID
	userplan.PlanID = userPlanCreate.PlanID
	userplan.Active = userPlanCreate.Active == "true"
	startdate, _ := time.Parse("2006-01-02 15:04", userPlanCreate.StartDate)
	userplan.StartDate = startdate
	expirydate, _ := time.Parse("2006-01-02 15:04", userPlanCreate.ExpiryDate)
	userplan.ExpiryDate = expirydate

	if err := userplan.Create(); err != nil {
		logger.Log(logger.Fields{
			Loc:  "/userplans/create - CreateUserPlan()",
			Code: errors.InternalServerError.Code,
			Extra: map[string]interface{}{
				"UserPlan": userplan,
			},
			Err: err.Error(),
		})
		c.AbortWithStatusJSON(errors.InternalServerError.Status, errors.InternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"data": gin.H{
			"userplan": userplan,
		},
	})
}

// DeleteUserPlan deletes a given users userplan. It will not delete a userplan fully however,
// it will just set a DeletedAt datetime on the record
func DeleteUserPlan(c *gin.Context) {
	userID, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	var userplan models.UserPlan
	userplan.UserID = uint(userID)
	if err := userplan.Find(); err != nil {
		logger.Log(logger.Fields{
			Loc:   "/userplans/delete/:id - DeleteUserPlan()",
			Code:  errors.UserPlanNotFound.Code,
			Extra: map[string]interface{}{"UserPlanID": c.Param("id")},
			Err:   err.Error(),
		})
		c.AbortWithStatusJSON(errors.UserPlanNotFound.Status, errors.UserPlanNotFound)
		return
	}

	if err := userplan.Delete(); err != nil {
		logger.Log(logger.Fields{
			Loc:   "/userplans/delete/:id - DeleteUserPlan()",
			Code:  errors.PlanNotFound.Code,
			Extra: map[string]interface{}{"UserPlanID": c.Param("id")},
			Err:   err.Error(),
		})
		c.AbortWithStatusJSON(errors.InternalServerError.Status, errors.InternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 200,
	})
}

// UpdateUserPlan updates an existing plan
func UpdateUserPlan(c *gin.Context) {
	userPlanID, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	var userplan models.UserPlan
	userplan.UserID = uint(userPlanID)

	type UserPlanUpdates struct {
		Active     string `json:"active" binding:"required"`
		StartDate  string `json:"start_date" binding:"required"`
		ExpiryDate string `json:"expiry_date" binding:"required"`
	}
	userPlanUdates := UserPlanUpdates{}

	if err := userplan.Find(); err != nil {
		logger.Log(logger.Fields{
			Loc:   "/userplans/update/:id - UpdateUserPlan()",
			Code:  errors.UserPlanNotFound.Code,
			Extra: map[string]interface{}{"UserID": c.Param("id")},
			Err:   err.Error(),
		})
		c.AbortWithStatusJSON(errors.UserPlanNotFound.Status, errors.UserPlanNotFound)
		return
	}

	if err := c.BindJSON(&userPlanUdates); err != nil {
		logger.Log(logger.Fields{
			Loc:   "/userplans/update/:id - UpdateUserPlan()",
			Code:  errors.InvalidForm.Code,
			Extra: map[string]interface{}{"UserPlanID": c.Param("id")},
			Err:   err.Error(),
		})
		c.AbortWithStatusJSON(errors.InvalidForm.Status, errors.InvalidForm)
		return
	}

	userplan.Active = userPlanUdates.Active == "true"
	startdate, _ := time.Parse("2006-01-02 15:04", userPlanUdates.StartDate)
	userplan.StartDate = startdate
	expirydate, _ := time.Parse("2006-01-02 15:04", userPlanUdates.ExpiryDate)
	userplan.ExpiryDate = expirydate
	if err := userplan.Save(); err != nil {
		logger.Log(logger.Fields{
			Loc:   "/userplans/update/:id - UpdateUserPlan()",
			Code:  errors.InternalServerError.Code,
			Extra: map[string]interface{}{"UserPlanID": c.Param("id")},
			Err:   err.Error(),
		})
		c.AbortWithStatusJSON(errors.InternalServerError.Status, errors.InternalServerError)
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 200,
	})
}

// AllUserPlans returns an array of all available plans
func AllUserPlans(c *gin.Context) {
	var plans models.AllUserPlans

	if err := plans.FindAll(); err != nil {
		logger.Log(logger.Fields{
			Loc:  "/userplans - AllUserPlans()",
			Code: errors.InternalServerError.Code,
			Err:  err.Error(),
		})
		c.AbortWithStatusJSON(errors.InternalServerError.Status, errors.InternalServerError)
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"data": gin.H{
			"userplans": plans,
		},
	})
}
