package models

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

// Plan holds the detilas for a given vpn plan on offer
type Plan struct {
	BaseModel
	Name           string `json:"name" binding:"required"`
	Type           string `json:"type" binding:"required"`
	DurationHours  *int   `json:"duration_hours" binding:"exists"`
	DurationDays   *int   `json:"duration_days" binding:"exists"`
	DurationMonths *int   `json:"duration_months" binding:"exists"`
}

// BeforeCreate sets the CreatedAt column to the current time
func (plan *Plan) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedAt", time.Now())

	return nil
}

// BeforeUpdate sets the UpdatedAt column to the current time
func (plan *Plan) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("UpdatedAt", time.Now())
	return nil
}

// BeforeCreate sets the CreatedAt column to the current time
func (plan *Plan) String() string {
	return fmt.Sprintf(
		"ID: %d, Name: %s, Type: %s, DurationHours: %d, DurationDays: %d, DurationMonths: %d",
		plan.ID,
		plan.Name,
		plan.Type,
		*plan.DurationHours,
		*plan.DurationDays,
		*plan.DurationMonths,
	)
}
