package validations

import (
	"errors"
	"log"
	"regexp"

	"github.com/go-playground/validator/v10"
)

// ^09([0-9]{9})$
func IranainMobileNumberValidator(fld validator.FieldLevel) bool {
	value, ok := fld.Field().Interface().(string)
	if !ok {
		return false
	}
	res, err := regexp.MatchString(`^09([0-9]{9})$`, value)
	if err != nil {
		log.Print(err.Error())
	}
	return res
}

func PasswoordValidator(fld validator.FieldLevel) bool {
	value, ok := fld.Field().Interface().(string)
	if !ok {
		return false
	}
	res, err := regexp.MatchString(`^[a-zA-Z0-9]{10,20}$`, value)
	if err != nil {
		log.Print(err.Error())
	}
	return res
}

func IdValidator(fld validator.FieldLevel) bool {
	value, ok := fld.Field().Interface().(string)
	if !ok {
		return false
	}
	res, err := regexp.MatchString(`^[a-zA-Z0-9_]{3,15}$`, value)
	if err != nil {
		log.Print(err.Error())
	}
	return res
}

type ValidationError struct {
	Property string `json:"property"`
	Tag      string `json:"tag"`
	Value    string `json:"value"`
	Message  string `json:"message"`
}


func GetValidationErrors(err error) *[]ValidationError {
	var validationErrors []ValidationError
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		for _, err := range err.(validator.ValidationErrors) {
			var el ValidationError
			el.Property = err.Field()
			el.Tag = err.Tag()
			el.Value = err.Param()
			validationErrors = append(validationErrors, el)
		}
		return &validationErrors
	}
	return nil
}