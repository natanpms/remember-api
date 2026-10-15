package main

import (
	"log"

	"gin-quickstart/db"
	"gin-quickstart/routes"

	"github.com/joho/godotenv"
)

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
