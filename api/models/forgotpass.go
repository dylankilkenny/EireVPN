package models

import (
	"time"

	"github.com/satori/go.uuid"
	"github.com/jinzhu/gorm"
)

// ForgotPassword contains the password access token with a one to one mapping
// to the user
type ForgotPassword struct {
	BaseModel
	UserID uint 
	Token  string `json:"token"`
}

func (fp *ForgotPassword) Find() error {
	if err := db().Where(&fp).First(&fp).Error; err != nil {
		return err
	}
	return nil
}

func (fp *ForgotPassword) Create() error {
	if err := db().Create(&fp).Error; err != nil {
		return err
	}
	return nil
}

func (fp *ForgotPassword) Save() error {
	if err := db().Save(&fp).Error; err != nil {
		return err
	}
	return nil
}

func (fp *ForgotPassword) Delete() error {
	if err := db().Delete(&fp).Error; err != nil {
		return err
	}
	return nil
}

// BeforeCreate sets the CreatedAt column to the current time
func (fp *ForgotPassword) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedAt", time.Now())
	token := uuid.NewV4()
	scope.SetColumn("Token", token.String())
	return nil
}

// BeforeUpdate sets the UpdatedAt column to the current time
func (fp *ForgotPassword) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("UpdatedAt", time.Now())
	return nil
}
