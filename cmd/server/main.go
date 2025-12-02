package main

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func main() {
	validate = validator.New()
	fmt.Println("Validator initialized")
}
