package validation

import (
	"github.com/amirmalekian/Mashinchi/src/common"
	"github.com/go-playground/validator/v10"
)

func IranianPhoneNumberValidator(fld validator.FieldLevel) bool {

	value, ok := fld.Field().Interface().(string)
	if !ok {
		return false
	}

	return common.IranianPhoneNumberValidate(value)
}
