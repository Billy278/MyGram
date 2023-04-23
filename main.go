package main

import (
	"log"

	"github.com/Billy278/MyGram/server"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	validate := validator.New()
	server.NewServer(validate)

}
