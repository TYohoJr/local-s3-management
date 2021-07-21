package main

import (
	"s3manager/controller"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	godotenv.Load()
	// Initialize server and start listening
	s := controller.Server{}
	s.Initialize()
}
