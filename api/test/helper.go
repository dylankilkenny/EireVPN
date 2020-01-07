package test

import (
	"eirevpn/api/config"
	"eirevpn/api/db"
	"eirevpn/api/errors"
	"eirevpn/api/models"
	"eirevpn/api/util/jwt"
	"encoding/json"
	"fmt"
	"log"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" //db
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

var err error
var dbInstance *gorm.DB
var r *gin.Engine

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

	conf := config.GetConfig()
	conf.DB.Port = 5431
	conf.DB.User = "eirevpn_test"
	conf.DB.Password = "eirevpn_test"
	conf.DB.Host = "localhost"
	conf.DB.Database = "eirevpn_test"

	db.Init(conf, false, models.Get())
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

// CreateAdminUser adds a new user to the db and returns the object
func CreateAdminUser() *models.User {
	user := models.User{
		FirstName: "Dylan",
		LastName:  "Kilkenny",
		Email:     "email@email.com",
		Password:  "password",
		Type:      models.UserTypeAdmin}
	err := dbInstance.Create(&user).Error
	if err != nil {
		fmt.Println("CreatUser() - ", err)
	}
	return &user
}

// CreatePlan creates a new plan record in the db
func CreatePlan() *models.Plan {
	plan := models.Plan{
		Name:          "test_plan",
		Amount:        100,
		Interval:      "month",
		IntervalCount: int64(1),
		Currency:      "EUR",
	}
	err := dbInstance.Create(&plan).Error
	if err != nil {
		fmt.Println("CreatePlan() - ", err)
	}
	return &plan
}

// CreateUserPlan creates a new user plan record in the db
func CreateUserPlan(planID, userID uint, active bool) *models.UserPlan {
	userPlan := models.UserPlan{
		UserID:     userID,
		PlanID:     planID,
		Active:     active,
		StartDate:  time.Now(),
		ExpiryDate: time.Now().Add(time.Hour),
	}
	err := dbInstance.Create(&userPlan).Error
	if err != nil {
		fmt.Println("CreateUserPlan() - ", err)
	}
	return &userPlan
}

// CreateServer creates a new server record in the db
func CreateServer() *models.Server {
	server := models.Server{
		Country:     "Ireland",
		CountryCode: "IE",
		Type:        models.ServerTypeProxy,
		IP:          "127.0.0.1",
		Port:        8080,
	}
	err := dbInstance.Create(&server).Error
	if err != nil {
		fmt.Println("CreateServer() - ", err)
	}
	return &server
}

// CreateCleanDB drops exisitng tables and recreates them
func CreateCleanDB() {
	dbInstance.DropTableIfExists(&models.User{})
	dbInstance.DropTableIfExists(&models.Plan{})
	dbInstance.DropTableIfExists(&models.UserAppSession{})
	dbInstance.DropTableIfExists(&models.Server{})
	dbInstance.DropTableIfExists(&models.UserPlan{})

	if !dbInstance.HasTable(&models.User{}) {
		dbInstance.CreateTable(&models.User{})
	}

	if !dbInstance.HasTable(&models.Plan{}) {
		dbInstance.CreateTable(&models.Plan{})
	}

	if !dbInstance.HasTable(&models.UserAppSession{}) {
		dbInstance.CreateTable(&models.UserAppSession{})
	}

	if !dbInstance.HasTable(&models.Server{}) {
		dbInstance.CreateTable(&models.Server{})
	}

	if !dbInstance.HasTable(&models.UserPlan{}) {
		dbInstance.CreateTable(&models.UserPlan{})
	}
}

// DropPlanTable dros the plan table from the db
func DropPlanTable() {
	dbInstance.DropTableIfExists(&models.Plan{})
}

// DropServerTable dros the server table from the db
func DropServerTable() {
	dbInstance.DropTableIfExists(&models.Server{})
}

// GetToken fetches a jwt token for the given user
func GetToken(u *models.User) (authToken, refreshToken, csrfToken string) {
	var usersession models.UserAppSession
	usersession.UserID = u.ID
	dbInstance.Create(&usersession)
	authToken, refreshToken, csrfToken, err := jwt.Tokens(usersession)
	if err != nil {
		//TODO: add internal server error response here
		fmt.Printf("Error creating auth token for user ")
	}
	return
}

// DeleteIdentifier removes the users session identifier
func DeleteIdentifier(u *models.User) {
	var usersession models.UserAppSession
	dbInstance.Where("user_id = ?", u.ID).First(&usersession)
	dbInstance.Delete(&usersession)
}
