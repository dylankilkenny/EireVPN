package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// UserPlan contains the details of which plans each user is signed up for
type UserPlan struct {
	BaseModel
	UserID        uint      `json:"user_id"`
	PlaneID       uint      `json:"plan_id"`
	StartDateTime time.Time `json:"start_datetime"`
	EndDateTime   time.Time `json:"end_datetime"`
	Status        int       `json:"status"`
}

// BeforeCreate sets the CreatedAt column to the current time
func (userplan *UserPlan) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedAt", time.Now())

	return nil
}

// BeforeUpdate sets the UpdatedAt column to the current time
func (userplan *UserPlan) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("UpdatedAt", "check")
	return nil
}
