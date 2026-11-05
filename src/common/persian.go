package common

import (
	"regexp"

	"github.com/amirmalekian/Mashinchi/src/config"
	"github.com/amirmalekian/Mashinchi/src/pkg/logging"
)

var logger = logging.NewLogger(config.GetConfig())

const iranianPhoneNumberPattern string = `^09(1[0-9]|2[0-2]|3[0-9]|9[0-9])[0-9]{7}$`

func IranianPhoneNumberValidate(phoneNumber string) bool {
	res, err := regexp.MatchString(iranianPhoneNumberPattern, phoneNumber)
	if err != nil {
		logger.Error(logging.Validation, logging.PhoneValidation, err.Error(), nil)
	}
	return res
}
