package models

import (
	"log"
	"time"

	uuid "github.com/gofrs/uuid"
	"github.com/jinzhu/gorm"
)

type Plan struct {
	BaseModel
	Name           string `json:"name"`
	Type           string `json:"type"`
	DurationHours  int    `json:"duration_hours"`
	DurationDays   int    `json:"duration_days"`
	DurationMonths int    `json:"duration_months"`
	Password       string `json:"password" binding:"required"`
}

func (plan *Plan) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedAt", time.Now())
	uuid, err := uuid.NewV4()
	if err != nil {
		log.Println("Plan{} - UUID error")
		panic(err)
	}
	scope.SetColumn("ID", uuid.String())
	return nil
}

func (plan *Plan) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("UpdatedAt", "check")
	return nil
}
