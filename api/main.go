package main

import (
	"eirevpn/api/config"
	"eirevpn/api/integrations"
	"eirevpn/api/logger"
	"eirevpn/api/models"
	"eirevpn/api/router"
	"os"
	"path/filepath"

	"eirevpn/api/db"
)

func main() {
	debugMode := false
	logging := true

	appPath, _ := os.Getwd()
	filename, _ := filepath.Abs(appPath + "/config.yaml")
	config.Init(filename)
	conf := config.GetConfig()

	integrations.Init()

	db.Init(conf, debugMode, models.Get())

	logger.Init(logging)

	r := router.Init(logging)

	r.Run(":" + conf.App.Port)
}
