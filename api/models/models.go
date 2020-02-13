package models

import (
	database "eirevpn/api/db"

	"github.com/jinzhu/gorm"
)

var db_connection *gorm.DB

func db() *gorm.DB {
	if db_connection == nil {
		db_connection = database.GetDB()
	}
	return db_connection
}

// Get returns all models
func Get() []interface{} {
	return []interface{}{
		&Plan{},
		&User{},
		&UserPlan{},
		&UserAppSession{},
		&Cart{},
		&Server{},
		&EmailToken{},
	}
}
