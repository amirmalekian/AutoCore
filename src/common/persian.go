package common

import (
	"log"
	"regexp"
)

const iranianPhoneNumberPattern string = `^09(1[0-9]|2[0-2]|3[0-9]|9[0-9])[0-9]{7}$`

func IranianPhoneNumberValidate(phoneNumber string) bool {
	res, err := regexp.MatchString(iranianPhoneNumberPattern, phoneNumber)
	if err != nil {
		log.Print(err.Error())
	}
	return res
}
