package models

import (
	"time"

	uuid "github.com/gofrs/uuid"
)

type BaseModel struct {
	ID        uuid.UUID  `gorm:"primary_key;column:id" json:"id"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `sql:"index" json:"deletedAt,omitempty"`
}
