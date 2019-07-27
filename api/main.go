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

	// if value, ok := os.LookupEnv(key); !ok {
	// 	log.Fatal("main() -> Error loading .env file")
	// }
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println(err)
		log.Fatal("main() -> Error loading .env file")
	}

	db.Init(debugMode)
	db := db.GetDB()

	r := router.SetupRouter(db, logging)
	r.Run(":" + os.Getenv("PORT"))
}
