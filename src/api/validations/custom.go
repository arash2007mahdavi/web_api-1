package validations

import (
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
