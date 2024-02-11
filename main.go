package main

import (
	"fmt"

	"github.com/vansh845/go-microservices/application"
)

func main() {
	app := application.New()

	err := app.Start()
	if err != nil {
		fmt.Println(err)
	}
}
