package models

import (
	"log"
	"time"

	uuid "github.com/gofrs/uuid"
	"github.com/jinzhu/gorm"
)

type UserPlan struct {
	BaseModel
	UserID        uuid.UUID `json:"user_id"`
	PlaneID       uuid.UUID `json:"plan_id"`
	StartDateTime time.Time `json:"start_datetime"`
	EndDateTime   time.Time `json:"end_datetime"`
	Status        int       `json:"status"`
}

func (userplan *UserPlan) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedAt", time.Now())
	uuid, err := uuid.NewV4()
	if err != nil {
		log.Println("Plan{} - UUID error")
		panic(err)
	}
	scope.SetColumn("ID", uuid.String())
	return nil
}

func (userplan *UserPlan) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("UpdatedAt", "check")
	return nil
}
