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
// @host localhost:8080
// @BasePath /

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and the JWT token.
func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	conn, err := db.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	router := routes.ExecuteRoutes(conn)
	router.Run() // listens on 0.0.0.0:8080 by default
}
