package main

import (
	"eirevpn/api/router"
	"os"

	"fmt"
	"log"

	"eirevpn/api/db"

	"github.com/joho/godotenv"
)

func main() {
	debugMode := false
	logging := true

	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println(err)
		log.Fatal("main() -> Error loading .env file")
	}
	config := db.DbConfig{}
	config.Load()
	db.Init(config, debugMode)

	r := router.SetupRouter(logging)
	r.Run(":" + os.Getenv("PORT"))
}
