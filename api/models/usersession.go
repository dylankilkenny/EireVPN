package models

import (
	"eirevpn/api/util/random"
	"time"

	"github.com/jinzhu/gorm"
)

// UserSession contains the details of which plans each user is signed up for
type UserSession struct {
	BaseModel
	UserID     uint   `json:"user_id"`
	Identifier string `json:"indentifier"`
}

// BeforeCreate sets the CreatedAt column to the current time
func (userplan *UserSession) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedAt", time.Now())
	if identifier, err := random.GenerateRandomString(64); err == nil {
		scope.SetColumn("Identifier", identifier)

	}
	return nil
}

// BeforeUpdate sets the UpdatedAt column to the current time
func (userplan *UserSession) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("UpdatedAt", "check")
	return nil
}
