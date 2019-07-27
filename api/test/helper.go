package test

import (
	"fmt"
	"log"

	"eirevpn/api/models"
	"eirevpn/api/util/jwt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" //db
	_ "github.com/lib/pq"
)

var db *gorm.DB
var err error

func InitDB() *gorm.DB {

	userDb := "dylankilkenny"
	password := ""
	host := "localhost"
	port := "5432"
	dbname := "test"

	dbinfo := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		userDb,
		password,
		host,
		port,
		dbname,
	)

	db, err = gorm.Open("postgres", dbinfo)

	if err != nil {
		log.Println("Failed to connect to testing database")
		panic(err)
	}
	log.Println("Testing Database connected")

	CreateCleanDB()

	return db
}

func CreateUser() *models.User {
	user := models.User{FirstName: "Dylan", LastName: "Kilkenny", Email: "email@email.com", Password: "password"}
	db.Create(&user)
	return &user
}

func CreateCleanDB() {
	db.DropTableIfExists(&models.User{})

	if !db.HasTable(&models.User{}) {
		db.CreateTable(&models.User{})
	}

}

func GetToken(u *models.User) (token string) {
	token, err := jwt.Token(u.ID.String())
	if err != nil {
		fmt.Printf("Error creating token for user %v", u.ID)
	}
	return
}