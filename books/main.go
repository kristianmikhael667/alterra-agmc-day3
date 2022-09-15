package main

import (
	"log"
	"main/routes"

	"github.com/joho/godotenv"
)

//----------
// Handlers
//----------

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	e := routes.New()
	e.Logger.Fatal(e.Start(":8080"))
}
