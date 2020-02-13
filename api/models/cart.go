package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Cart contains the details of which plans each user is trying to purchase
type Cart struct {
	BaseModel
	UserID uint `json:"user_id"`
	PlanID uint `json:"plan_id"`
}

func (c *Cart) Find() error {
	if err := db().Where(&c).First(&c).Error; err != nil {
		return err
	}
	return nil
}

func (c *Cart) Save() error {
	if err := db().Save(&c).Error; err != nil {
		return err
	}
	return nil
}

func (c *Cart) Delete() error {
	if err := db().Delete(&c).Error; err != nil {
		return err
	}
	return nil
}

// BeforeCreate sets the CreatedAt column to the current time
func (c *Cart) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedAt", time.Now())

	return nil
}

// BeforeUpdate sets the UpdatedAt column to the current time
func (c *Cart) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("UpdatedAt", time.Now())
	return nil
}
