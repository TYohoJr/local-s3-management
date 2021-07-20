package main

import (
	"s3manager/controller"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	s := controller.Server{}
	s.Initialize()
}
