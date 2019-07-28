package models

import (
	"time"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

// User contains the users details
type User struct {
	BaseModel
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password" binding:"required"`
}

// BeforeCreate sets the CreatedAt column to the current time
// and encrypts the users password
func (user *User) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedAt", time.Now())

	if pw, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost); err == nil {
		scope.SetColumn("Password", string(pw))
	}

	return nil
}

// BeforeUpdate sets the UpdatedAt column to the current time
func (user *User) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("UpdatedAt", time.Now())
	return nil
}
