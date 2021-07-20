package main

import (
	"fmt"
	"s3manager/controller"
)

func main() {
	fmt.Println("hello")
	s := controller.Server{}
	s.Initialize()
}
