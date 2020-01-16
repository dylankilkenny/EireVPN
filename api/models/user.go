package models

import (
	"eirevpn/api/integrations/stripe"
	"time"

	stripego "github.com/stripe/stripe-go"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type UserType string
type AllUsers []User

var (
	UserTypeNormal UserType = "normal"
	UserTypeAdmin  UserType = "admin"
)

// User contains the users details
type User struct {
	BaseModel
	FirstName        string   `json:"firstname"`
	LastName         string   `json:"lastname"`
	Email            string   `json:"email" binding:"required"`
	Password         string   `json:"password" binding:"required"`
	StripeCustomerID string   `json:"stripe_customer_id"`
	Type             UserType `json:"type"`
}

func (u *User) Find() error {
	if err := db().Where(u).First(&u).Error; err != nil {
		return err
	}
	return nil
}

func (u *User) Create() error {
	if err := db().Create(&u).Error; err != nil {
		return err
	}
	return nil
}

func (u *User) Save() error {
	if err := db().Save(&u).Error; err != nil {
		return err
	}
	return nil
}

func (u *User) CreateStripeCustomer() (*stripego.Customer, error) {
	return stripe.CreateCustomer(u.Email, u.FirstName, u.LastName, u.ID)
}

func (au *AllUsers) FindAll() error {
	if err := db().Find(&au).Error; err != nil {
		return err
	}
	return nil
}

// BeforeCreate sets the CreatedAt column to the current time
// and encrypts the users password
func (u *User) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedAt", time.Now())
	if pw, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost); err == nil {
		scope.SetColumn("Password", string(pw))
	}
	return nil
}

// BeforeUpdate sets the UpdatedAt column to the current time
func (u *User) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("UpdatedAt", time.Now())
	return nil
}
