package dto

type GetOtpRequest struct {
	PhoneNumber string `json:"phoneNumber" binding:"required,phone,min=11,max=11"`
}
