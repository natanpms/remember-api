package main

import (
	"log"

	"github.com/natanpms/remember-api/db"
	"github.com/natanpms/remember-api/routes"

	"github.com/joho/godotenv"
)

// @title Remember API
// @version 1.0
// @description API for the Remember app.
// @host localhost:8000
// @BasePath /
func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	conn, err := db.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	router := routes.ExecuteRoutes()
	router.Run() // listens on 0.0.0.0:8080 by default
}
