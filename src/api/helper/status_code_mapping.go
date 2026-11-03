package helper

import (
	"net/http"

	"github.com/amirmalekian/Mashinchi/src/pkg/service_errors"
)

var StatusCodeMapping = map[string]int{

	// OTP
	service_errors.OtpExists:  409,
	service_errors.OtpUsed:    409,
	service_errors.OtpInValid: 400,
}

func TranslateErrorToStatusCode(err error) int {
	if value, ok := StatusCodeMapping[err.Error()]; !ok {
		return http.StatusInternalServerError
	} else {
		return value
	}
}
