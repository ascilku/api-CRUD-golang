package user

import "github.com/go-playground/validator/v10"

type User struct {
	ID       int
	Name     string
	Email    string
	Password string
	File     string
}

func ErrorValidation(err error) []string {
	var errors []string
	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, e.Error())
	}
	return errors
}
