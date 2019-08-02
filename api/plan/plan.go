package plan

import (
	"eirevpn/api/db"
	"eirevpn/api/errors"
	"eirevpn/api/logger"
	"eirevpn/api/models"
	"net/http"

	logrus "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

// Plan fetches a plan by ID
func Plan(c *gin.Context) {
	id := c.Param("id")
	db := db.GetDB()
	var p models.Plan

	if err := db.Where("id = ?", id).First(&p).Error; err != nil {
		logger.Log(logrus.Fields{
			"Loc:":    "/plan/:id - Plan()",
			"Code:":   errors.PlanNotFound.Code,
			"PlanID:": id,
		}, err.Error())
		c.AbortWithStatusJSON(errors.PlanNotFound.Status, errors.PlanNotFound)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"errors": make([]string, 0),
		"data": gin.H{
			"plan": p,
		},
	})

}

// CreatePlan creates a new plan
func CreatePlan(c *gin.Context) {
	db := db.GetDB()
	var p models.Plan

	if err := c.BindJSON(&p); err != nil {
		logger.Log(logrus.Fields{
			"Loc:":  "/plan - CreatePlan()",
			"Code:": errors.InvalidForm.Code,
		}, err.Error())
		c.AbortWithStatusJSON(errors.InvalidForm.Status, errors.InvalidForm)
		return
	}

	if err := db.Create(&p).Error; err != nil {
		logger.Log(logrus.Fields{
			"Loc:":  "/plan - CreatePlan()",
			"Code:": errors.InternalServerError.Code,
		}, err.Error())
		c.AbortWithStatusJSON(errors.InternalServerError.Status, errors.InternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"errors": make([]string, 0),
		"data": gin.H{
			"plan": p,
		},
	})
}

// DeletePlan deletes a given plan. It will not delete a plan fully however,
// it will just set a DeletedAt datetime on the record
func DeletePlan(c *gin.Context) {
	id := c.Param("id")
	db := db.GetDB()
	var p models.Plan

	if err := db.Where("id = ?", id).First(&p).Error; err != nil {
		logger.Log(logrus.Fields{
			"Loc:":   "/plan/:id - DeletePlan()",
			"Code:":  errors.PlanNotFound.Code,
			"PlanID": id,
		}, err.Error())
		c.AbortWithStatusJSON(errors.PlanNotFound.Status, errors.PlanNotFound)
		return
	}

	if err := db.Delete(&p).Error; err != nil {
		logger.Log(logrus.Fields{
			"Loc:":   "/plan/:id - DeletePlan()",
			"Code:":  errors.InternalServerError.Code,
			"PlanID": id,
		}, err.Error())
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
	db := db.GetDB()
	var p models.Plan

	if err := c.BindJSON(&p); err != nil {
		logger.Log(logrus.Fields{
			"Loc:":   "/plan - UpdatePlan()",
			"Code:":  errors.InvalidForm.Code,
			"PlanID": p.ID,
		}, err.Error())
		c.AbortWithStatusJSON(errors.InvalidForm.Status, errors.InvalidForm)
		return
	}

	if err := db.Save(&p).Error; err != nil {
		logger.Log(logrus.Fields{
			"Loc:":   "/plan - UpdatePlan()",
			"Code:":  errors.InternalServerError.Code,
			"PlanID": p.ID,
		}, err.Error())
		c.AbortWithStatusJSON(errors.InternalServerError.Status, errors.InternalServerError)
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"errors": make([]string, 0),
		"data":   make([]string, 0),
	})
}

// AllPlans returns an array of all available plans
func AllPlans(c *gin.Context) {
	db := db.GetDB()
	var plans []models.Plan

	if err := db.Find(&plans).Error; err != nil {
		logger.Log(logrus.Fields{
			"Loc:":  "/plans - AllPlans()",
			"Code:": errors.InternalServerError.Code,
		}, err.Error())
		c.AbortWithStatusJSON(errors.InternalServerError.Status, errors.InternalServerError)
	}

	if len(plans) == 0 {
		c.AbortWithStatusJSON(errors.NoPlansFound.Status, errors.NoPlansFound)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"errors": make([]string, 0),
		"data": gin.H{
			"plans": plans,
		},
	})
}
