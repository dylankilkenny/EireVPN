package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type AllServers []Server

type ServerType string

var (
	ServerTypeProxy ServerType = "Proxy"
	ServerTypeVPN   ServerType = "VPN"
)

// Cart contains the details of which plans each user is signed trying to purchase
type Server struct {
	BaseModel
	Country     string     `form:"country" json:"country" binding:"required"`
	CountryCode string     `form:"country_code" json:"country_code" binding:"required"`
	Type        ServerType `form:"type" json:"type"` // binding:"required" removed for time being
	IP          string     `form:"ip" json:"ip" binding:"required"`
	Port        int        `form:"port" json:"port" binding:"required"`
	Username    string     `form:"username" json:"username" binding:"required"`
	Password    string     `form:"password" json:"password" binding:"required"`
	ImagePath   string     `form:"image_path" json:"image_path"`
}

func (s *Server) Find() error {
	if err := db().Where(&s).First(&s).Error; err != nil {
		return err
	}
	return nil
}

func (s *Server) Create() error {
	if err := db().Create(&s).Error; err != nil {
		return err
	}
	return nil
}

func (s *Server) Save() error {
	if err := db().Save(&s).Error; err != nil {
		return err
	}
	return nil
}

func (s *Server) Delete() error {
	if err := db().Delete(&s).Error; err != nil {
		return err
	}
	return nil
}

func (as *AllServers) FindAll() error {
	if err := db().Find(&as).Error; err != nil {
		return err
	}
	return nil
}

// BeforeCreate sets the CreatedAt column to the current time
func (s *Server) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedAt", time.Now())

	return nil
}

// Beforecdate sets the cdatedAt column to the current time
func (s *Server) Beforecdate(scope *gorm.Scope) error {
	scope.SetColumn("cdatedAt", "check")
	return nil
}
