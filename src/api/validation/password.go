package validation

import (
	"github.com/amirmalekian/Mashinchi/src/common"
	"github.com/amirmalekian/Mashinchi/src/config"
	"github.com/go-playground/validator/v10"
)

func PasswordValidator(cfg *config.Config, fld validator.FieldLevel) bool {
	value, ok := fld.Field().Interface().(string)
	if !ok {
		fld.Param()
		return false
	}

	return common.CheckPassword(cfg, value)
}
