package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type AllConnections []Connection

// Connections contains the email confirmation tokens with a one to one mapping
// to the user
type Connection struct {
	BaseModel
	UserID        uint   `json:"user_id"`
	ServerID      uint   `json:"server_id"`
	ServerCountry string `json:"server_country"`
}

func (c *Connection) Find() error {
	if err := db().Where(&c).First(&c).Error; err != nil {
		return err
	}
	return nil
}

func (c *Connection) GetUser() (*User, error) {
	var user User
	if err := db().Model(&user).Related(&c).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (c *Connection) Create() error {
	if err := db().Create(&c).Error; err != nil {
		return err
	}
	return nil
}

func (c *Connection) Save() error {
	if err := db().Save(&c).Error; err != nil {
		return err
	}
	return nil
}

func (c *Connection) Delete() error {
	if err := db().Delete(&c).Error; err != nil {
		return err
	}
	return nil
}

func (ac *AllConnections) FindAll(offset int) error {
	limit := 20
	if err := db().Order("created_at desc").Limit(limit).Offset(offset).Find(&ac).Error; err != nil {
		return err
	}
	return nil
}

func (ac *AllConnections) Count() (*int, error) {
	var count int
	if err := db().Model(&AllConnections{}).Count(&count).Error; err != nil {
		return nil, err
	}
	return &count, nil
}

// BeforeCreate sets the CreatedAt column to the current time
func (c *Connection) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedAt", time.Now())
	return nil
}

// BeforeUpdate sets the UpdatedAt column to the current time
func (c *Connection) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("UpdatedAt", time.Now())
	return nil
}
