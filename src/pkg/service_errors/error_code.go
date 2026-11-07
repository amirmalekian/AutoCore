package service_errors

const (
	// Token
	UnExpectedError = "Expected Error"
	ClaimsNotFound  = "Claims Not Found"
	TokenRequired   = "Token Required"
	TokenExpired    = "Token Expired"
	TokenInvalid    = "Token Invalid"

	// OTP
	OtpExists  = "Otp Exists"
	OtpUsed    = "Otp Used"
	OtpInValid = "Otp Invalid"

	// User
	EmailExists    = "Email Exists"
	UsernameExists = "Username Exists"
)
