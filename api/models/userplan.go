package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type AllUserPlans []UserPlan

// UserPlan contains the details of which plans each user is signed up for
type UserPlan struct {
	BaseModel
	UserID     uint      `json:"user_id" binding:"required"`
	PlanID     uint      `json:"plan_id" binding:"required"`
	Active     bool      `json:"active" binding:"required"`
	StartDate  time.Time `json:"start_date" binding:"required"`
	ExpiryDate time.Time `json:"expiry_date" binding:"required"`
}

func (up *UserPlan) Find() error {
	if err := db().Where(&up).First(&up).Error; err != nil {
		return err
	}
	return nil
}

func (aup *AllUserPlans) FindAll() error {
	if err := db().Find(&aup).Error; err != nil {
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

func (up *UserPlan) Create() error {
	if err := up.DeleteAll(); err != nil {
		return err
	}
	if err := db().Create(&up).Error; err != nil {
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

// DeleteAll removes any existing user sessions
func (up *UserPlan) DeleteAll() error {
	if err := db().Delete(UserPlan{}, "user_id = ?", up.UserID).Error; err != nil {
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
