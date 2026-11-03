package dto

type GetOtpRequest struct {
	PhoneNumber string `json:"phoneNumber" binding:"required,phone,min=11,max=11"`
}

type TokenDetail struct {
	AccessToken            string `json:"accessToken"`
	RefreshToken           string `json:"refreshToken"`
	AccessTokenExpireTime  int64  `json:"accessTokenExpireTime"`
	RefreshTokenExpireTime int64  `json:"refreshTokenExpireTime"`
}
