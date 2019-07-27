package models

import (
	"log"
	"time"

	uuid "github.com/gofrs/uuid"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	BaseModel
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password" binding:"required"`
}

func (user *User) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedAt", time.Now())
	uuid, err := uuid.NewV4()
	if err != nil {
		log.Println("Uuid err")
		panic(err)
	}
	scope.SetColumn("ID", uuid.String())
	if pw, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost); err == nil {
		scope.SetColumn("Password", string(pw))
	}

	return nil
}

func (user *User) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("UpdatedAt", "check")
	return nil
}
