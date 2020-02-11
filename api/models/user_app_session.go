package models

import (
	"eirevpn/api/util/random"
	"time"

	"github.com/jinzhu/gorm"
)

// UserAppSession contains the users session identifier token
type UserAppSession struct {
	BaseModel
	UserID      uint   `json:"user_id"`
	Identifier  string `json:"indentifier"`
}

func (us *UserAppSession) Find() error {
	if err := db().Where(&us).First(&us).Error; err != nil {
		return err
	}
	return nil
}

// New adds a new user session and deletes all old sessions
func (us *UserAppSession) New(UserID uint) error {
	us.UserID = UserID
	if err := us.DeleteAll(); err != nil {
		return err
	}
	if err := us.Create(); err != nil {
		return err
	}
	return nil
}

// Create adds a new user sessions
func (us *UserAppSession) Create() error {
	if err := db().Create(&us).Error; err != nil {
		return err
	}
	return nil
}

// DeleteAll removes any existing user sessions
func (us *UserAppSession) DeleteAll() error {
	if err := db().Delete(UserAppSession{}, "user_id = ?", us.UserID).Error; err != nil {
		return err
	}
	return nil
}

// BeforeCreate sets the CreatedAt column to the current time
func (us *UserAppSession) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedAt", time.Now())
	if identifier, err := random.GenerateRandomString(64); err == nil {
		scope.SetColumn("Identifier", identifier)

	}
	return nil
}

// BeforeUpdate sets the UpdatedAt column to the current time
func (us *UserAppSession) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("UpdatedAt", "check")
	return nil
}
