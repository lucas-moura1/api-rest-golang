package validator

import (
	"errors"
	"fmt"
	"log"
	models "user-api/domain/model"

	"github.com/go-playground/validator/v10"
)

var v = validator.New()

func ValidateUserId(id string) error {
    err := v.Var(id, "required,hexadecimal")
    if err != nil {
        log.Print("Error: ", err)
        return err
    }

    return nil
}

func ValidateUser(u models.UserInput) error {
    err := v.Struct(u)
    if err == nil {
        return nil
    }

    log.Print("Error: ", err)
    errorMessage := handleErrorMessage(err)
    return errorMessage
}


func handleErrorMessage(err error) error {
    validationErr, _ := err.(validator.ValidationErrors)
    messageError := fmt.Sprintf("'%s' has a incorrect value: '%v'.", validationErr[0].Field(), validationErr[0].Value())

    return errors.New(messageError)
}
