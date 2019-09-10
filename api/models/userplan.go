package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// UserPlan contains the details of which plans each user is signed up for
type UserPlan struct {
	BaseModel
	UserID     uint      `json:"user_id"`
	PlanID     uint      `json:"plan_id"`
	Active     bool      `json:"active"`
	StartDate  time.Time `json:"start_date"`
	ExpiryDate time.Time `json:"expiry_date"`
}

func (up *UserPlan) Find() error {
	if err := db().Where(&up).First(&up).Error; err != nil {
		return err
	}
	return nil
}

func (up *UserPlan) Save() error {
	if err := db().Save(&up).Error; err != nil {
		return err
	}
	return nil
}

func (up *UserPlan) Delete() error {
	if err := db().Delete(&up).Error; err != nil {
		return err
	}
	return nil
}

// BeforeCreate sets the CreatedAt column to the current time
func (up *UserPlan) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedAt", time.Now())

	return nil
}

// BeforeUpdate sets the UpdatedAt column to the current time
func (up *UserPlan) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("UpdatedAt", "check")
	return nil
}
