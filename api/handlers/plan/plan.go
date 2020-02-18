package plan

import (
	"eirevpn/api/errors"
	"eirevpn/api/logger"
	"eirevpn/api/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Plan fetches a plan by ID
func Plan(c *gin.Context) {
	planID, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	var plan models.Plan
	plan.ID = uint(planID)
	if err := plan.Find(); err != nil {
		logger.Log(logger.Fields{
			Loc:   "/plan/:id - Plan()",
			Code:  errors.PlanNotFound.Code,
			Extra: map[string]interface{}{"PlanID": c.Param("id")},
			Err:   err.Error(),
		})
		c.AbortWithStatusJSON(errors.PlanNotFound.Status, errors.PlanNotFound)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"errors": make([]string, 0),
		"data": gin.H{
			"plan": plan,
		},
	})

}

// CreatePlan creates a new plan
func CreatePlan(c *gin.Context) {
	var plan models.Plan

	if err := c.BindJSON(&plan); err != nil {
		logger.Log(logger.Fields{
			Loc:  "/plans/create - CreatePlan()",
			Code: errors.InvalidForm.Code,
			Err:  err.Error(),
		})
		c.AbortWithStatusJSON(errors.InvalidForm.Status, errors.InvalidForm)
		return
	}

	if err := plan.Create(); err != nil {
		logger.Log(logger.Fields{
			Loc:  "/plans/create - CreatePlan()",
			Code: errors.InternalServerError.Code,
			Extra: map[string]interface{}{
				"Plan": plan,
			},
			Err: err.Error(),
		})
		c.AbortWithStatusJSON(errors.InternalServerError.Status, errors.InternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"errors": make([]string, 0),
		"data": gin.H{
			"plan": plan,
		},
	})
}

// DeletePlan deletes a given plan. It will not delete a plan fully however,
// it will just set a DeletedAt datetime on the record
func DeletePlan(c *gin.Context) {
	planID, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	var plan models.Plan
	plan.ID = uint(planID)
	if err := plan.Find(); err != nil {
		logger.Log(logger.Fields{
			Loc:   "/plans/delete/:id - DeletePlan()",
			Code:  errors.PlanNotFound.Code,
			Extra: map[string]interface{}{"PlanID": c.Param("id")},
			Err:   err.Error(),
		})
		c.AbortWithStatusJSON(errors.PlanNotFound.Status, errors.PlanNotFound)
		return
	}

	if err := plan.Delete(); err != nil {
		logger.Log(logger.Fields{
			Loc:   "/plans/delete/:id - DeletePlan()",
			Code:  errors.PlanNotFound.Code,
			Extra: map[string]interface{}{"PlanID": c.Param("id")},
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

// UpdatePlan updates an existing plan
func UpdatePlan(c *gin.Context) {
	planID, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	var plan models.Plan
	plan.ID = uint(planID)

	type PlanUpdates struct {
		Name string `json:"name" binding:"required"`
	}
	planUdates := PlanUpdates{}

	if err := plan.Find(); err != nil {
		logger.Log(logger.Fields{
			Loc:   "/plans/update/:id - UpdatePlan()",
			Code:  errors.PlanNotFound.Code,
			Extra: map[string]interface{}{"PlanID": c.Param("id")},
			Err:   err.Error(),
		})
		c.AbortWithStatusJSON(errors.PlanNotFound.Status, errors.PlanNotFound)
		return
	}

	if err := c.BindJSON(&planUdates); err != nil {
		logger.Log(logger.Fields{
			Loc:   "/plans/update/:id - UpdatePlan()",
			Code:  errors.InvalidForm.Code,
			Extra: map[string]interface{}{"PlanID": c.Param("id")},
			Err:   err.Error(),
		})
		c.AbortWithStatusJSON(errors.InvalidForm.Status, errors.InvalidForm)
		return
	}

	plan.Name = planUdates.Name
	if err := plan.Save(); err != nil {
		logger.Log(logger.Fields{
			Loc:   "/plans/update/:id - UpdatePlan()",
			Code:  errors.InternalServerError.Code,
			Extra: map[string]interface{}{"PlanID": plan.ID},
			Err:   err.Error(),
		})
		c.AbortWithStatusJSON(errors.InternalServerError.Status, errors.InternalServerError)
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"errors": make([]string, 0),
		"data":   make([]string, 0),
	})
}

// AllPlansPublic returns an array of all available plans only
// containing customer facing data
func AllPlansPublic(c *gin.Context) {
	var plans models.AllPlans

	if err := plans.FindAll(); err != nil {
		logger.Log(logger.Fields{
			Loc:  "/plans - AllPlans()",
			Code: errors.InternalServerError.Code,
			Err:  err.Error(),
		})
		c.AbortWithStatusJSON(errors.InternalServerError.Status, errors.InternalServerError)
	}

	publicPlanData := make([]map[string]interface{}, 0)
	for _, p := range plans {
		publicPlanData = append(publicPlanData,
			map[string]interface{}{
				"name":     p.Name,
				"amount":   p.Amount,
				"interval": p.Interval,
			})
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"data": gin.H{
			"plans": publicPlanData,
		},
	})
}

// AllPlans returns an array of all available plans
func AllPlans(c *gin.Context) {
	var plans models.AllPlans

	if err := plans.FindAll(); err != nil {
		logger.Log(logger.Fields{
			Loc:  "/plans - AllPlans()",
			Code: errors.InternalServerError.Code,
			Err:  err.Error(),
		})
		c.AbortWithStatusJSON(errors.InternalServerError.Status, errors.InternalServerError)
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"data": gin.H{
			"plans": plans,
		},
	})
}
