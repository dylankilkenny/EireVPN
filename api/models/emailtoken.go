package models

import (
	"time"

	"github.com/satori/go.uuid"
	"github.com/jinzhu/gorm"
)

// EmailConfirm contains the email confirmation tokens with a one to one mapping
// to the user
type EmailToken struct {
	BaseModel
	UserID uint 
	Token  string `json:"token"`
}

func (et *EmailToken) Find() error {
	if err := db().Where(&et).First(&et).Error; err != nil {
		return err
	}
	return nil
}

func (et *EmailToken) GetUser() (*User, error) {
	var user User
	if err := db().Model(&user).Related(&et).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (et *EmailToken) Create() error {
	if err := db().Create(&et).Error; err != nil {
		return err
	}
	return nil
}

func (et *EmailToken) Save() error {
	if err := db().Save(&et).Error; err != nil {
		return err
	}
	return nil
}

func (et *EmailToken) Delete() error {
	if err := db().Delete(&et).Error; err != nil {
		return err
	}
	return nil
}

// BeforeCreate sets the CreatedAt column to the current time
func (et *EmailToken) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedAt", time.Now())
	token := uuid.NewV4()
	scope.SetColumn("Token", token.String())
	return nil
}

// BeforeUpdate sets the UpdatedAt column to the current time
func (et *EmailToken) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("UpdatedAt", time.Now())
	return nil
}
