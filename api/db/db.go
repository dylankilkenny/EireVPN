package db

import (
	"eirevpn/api/config"
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

var db *gorm.DB
var err error

// Init creates a connection to postgres database and
// migrates any new models
func Init(config config.Config, debug bool, models []interface{}) {

	dbinfo := fmt.Sprintf("user=%s password=%s host=%s port=%v dbname=%s sslmode=disable",
		config.DB.User,
		config.DB.Password,
		config.DB.Host,
		config.DB.Port,
		config.DB.Database,
	)

	fmt.Println(dbinfo)

	db, err = gorm.Open("postgres", dbinfo)
	db.LogMode(debug)

	if err != nil {
		log.Println("Failed to connect to database")
		panic(err.Error())
	}
	log.Println("Database connected")

	for _, model := range models {
		if !db.HasTable(model) {
			err := db.CreateTable(model).Error
			if err != nil {
				log.Fatal(err)
			}
			log.Println("Table Created")
		}
		db.AutoMigrate(model)
	}
}

//GetDB ...
func GetDB() *gorm.DB {
	return db
}

func CloseDB() {
	db.Close()
}
