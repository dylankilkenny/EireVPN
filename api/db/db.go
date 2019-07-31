package db

import (
	"eirevpn/api/models"
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

var db *gorm.DB
var err error

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

// Init creates a connection to postgres database and
// migrates any new models
func Init(config DbConfig, debug bool) {

	dbinfo := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.Database,
	)

	fmt.Println(dbinfo)

	db, err = gorm.Open("postgres", dbinfo)
	db.LogMode(debug)

	if err != nil {
		log.Println("Failed to connect to database")
		panic(err)
	}
	log.Println("Database connected")
	if !db.HasTable(&models.User{}) {
		err := db.CreateTable(&models.User{})
		if err == nil {
			log.Println("Table Created")
		}
	}

	if !db.HasTable(&models.Plan{}) {
		err := db.CreateTable(&models.Plan{})
		if err == nil {
			log.Println("Table Created")
		}
	}

	if !db.HasTable(&models.UserPlan{}) {
		err := db.CreateTable(&models.UserPlan{})
		if err == nil {
			log.Println("Table Created")
		}
	}

	if !db.HasTable(&models.UserSession{}) {
		err := db.CreateTable(&models.UserSession{})
		if err == nil {
			log.Println("Table Created")
		}
	}

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Plan{})
	db.AutoMigrate(&models.UserPlan{})
	db.AutoMigrate(&models.UserSession{})
}

//GetDB ...
func GetDB() *gorm.DB {
	return db
}

func CloseDB() {
	db.Close()
}
