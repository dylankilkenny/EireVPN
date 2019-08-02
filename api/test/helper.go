package test

import (
	"eirevpn/api/db"
	"eirevpn/api/errors"
	"eirevpn/api/models"
	"eirevpn/api/util/jwt"
	"encoding/json"
	"fmt"
	"log"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" //db
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

var err error
var dbInstance *gorm.DB
var r *gin.Engine

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func assertCorrectStatus(t *testing.T, want, got int) {
	t.Helper()
	ok := assert.Equal(t, want, got)
	if !ok {
		t.Errorf("Status is not %v. Got %v", want, got)
	}
}

func bindError(resp *httptest.ResponseRecorder) errors.APIError {
	decoder := json.NewDecoder(resp.Body)
	var apiErr errors.APIError
	err := decoder.Decode(&apiErr)
	if err != nil {
		panic(err)
	}
	return apiErr
}

func assertCorrectCode(t *testing.T, want, got string) {
	t.Helper()
	ok := assert.Equal(t, want, got)
	if !ok {
		t.Errorf("Code is not %v. Got %v", want, got)
	}
}

// InitDB Creates a clean test database
func InitDB() {

	conf := db.DbConfig{}
	conf.Port = "5432"
	conf.User = "test"
	conf.Password = "test"
	conf.Host = "localhost"
	conf.Database = "eirevpn_test"

	db.Init(conf, false)
	dbInstance = db.GetDB()

	if err != nil {
		log.Println("Failed to connect to testing database")
		panic(err)
	}
	log.Println("Testing Database connected")

	CreateCleanDB()
}

// CreateUser adds a new user to the db and returns the object
func CreateUser() *models.User {
	user := models.User{FirstName: "Dylan", LastName: "Kilkenny", Email: "email@email.com", Password: "password"}
	err := dbInstance.Create(&user).Error
	if err != nil {
		fmt.Println("CreatUser() - ", err)
	}
	return &user
}

func CreatePlan() *models.Plan {
	intRef := func(i int) *int { return &i }
	plan := models.Plan{
		Name:           "test_plan",
		Type:           "test",
		DurationHours:  intRef(0),
		DurationDays:   intRef(0),
		DurationMonths: intRef(1),
	}
	err := dbInstance.Create(&plan).Error
	if err != nil {
		fmt.Println("CreatePlan() - ", err)
	}
	return &plan
}

// CreateCleanDB drops exisitng tables and recreates them
func CreateCleanDB() {
	dbInstance.DropTableIfExists(&models.User{})
	dbInstance.DropTableIfExists(&models.Plan{})
	dbInstance.DropTableIfExists(&models.UserSession{})

	if !dbInstance.HasTable(&models.User{}) {
		dbInstance.CreateTable(&models.User{})
	}

	if !dbInstance.HasTable(&models.Plan{}) {
		dbInstance.CreateTable(&models.Plan{})
	}

	if !dbInstance.HasTable(&models.UserSession{}) {
		dbInstance.CreateTable(&models.UserSession{})
	}
}

func DropPlanTable() {
	dbInstance.DropTableIfExists(&models.Plan{})
}

// GetToken fetches a jwt token for the given user
func GetToken(u *models.User) (authToken, refreshToken, csrfToken string) {
	var usersession models.UserSession
	usersession.UserID = u.ID
	dbInstance.Create(&usersession)
	authToken, refreshToken, csrfToken, err := jwt.Token(usersession)
	if err != nil {
		//TODO: add internal server error response here
		fmt.Printf("Error creating auth token for user ")
	}
	return
}

func DeleteIdentifier(u *models.User) {
	var usersession models.UserSession
	dbInstance.Where("user_id = ?", u.ID).First(&usersession)
	dbInstance.Delete(&usersession)
}
