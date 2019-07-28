package models

import (
	"time"
)

// BaseModel are commonly used fields for all models
type BaseModel struct {
	ID        uint       `gorm:"AUTO_INCREMENT;primary_key;column:id;" json:"id"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `sql:"index" json:"deletedAt,omitempty"`
}
