package plan

import (
	"eirevpn/api/db"
	"eirevpn/api/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Plan fetches a plan by ID
func Plan(c *gin.Context) {
	id := c.Param("id")
	db := db.GetDB()
	var p models.Plan

	if err := db.Where("id = ?", id).First(&p).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"errors": gin.H{
				"title":  "Plan Not Found",
				"detail": "No plan was found matching the queried id",
			},
			"data": make([]string, 0),
		})
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

	if err := db.Create(&p).Error; err != nil {
		fmt.Println("CreatePlan() -> ", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": 500,
			"errors": gin.H{
				"title":  "Internal Server Error",
				"detail": "An unkown error occured",
			},
			"data": make([]string, 0),
		})
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
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"errors": gin.H{
				"title":  "Plan Not Found",
				"detail": "No plan was found matching the queried id",
			},
			"data": make([]string, 0),
		})
		return
	}

	db.Delete(&p)

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

	if err := db.Save(&p).Error; err != nil {
		fmt.Println("UpdatePlan() -> ", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": 500,
			"errors": gin.H{
				"title":  "Internal Server Error",
				"detail": "An unkown error occured",
			},
			"data": make([]string, 0),
		})
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
		fmt.Println("AllPlans() -> ", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": 500,
			"errors": gin.H{
				"title":  "Internal Server Error",
				"detail": "An unkown error occured",
			},
			"data": make([]string, 0),
		})
	}

	if len(plans) == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"errors": gin.H{
				"title":  "No Plans Found",
				"detail": "There were no plans found",
			},
			"data": make([]string, 0),
		})
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
