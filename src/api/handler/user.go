package handler

import (
	"net/http"

	"github.com/amirmalekian/Mashinchi/src/api/dto"
	"github.com/amirmalekian/Mashinchi/src/api/helper"
	"github.com/amirmalekian/Mashinchi/src/config"
	"github.com/amirmalekian/Mashinchi/src/services"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service *services.UserService
}

func NewUserHandler(cfg *config.Config) *UserHandler {
	service := services.NewUserService(cfg)
	return &UserHandler{service: service}
}

// SendOtp 		godoc
// @Summary 	Send Otp to user
// @Description Send Otp to user
// @Tags 		Users
// @Accept  	json
// @Produce  	json
// @Param 		Request body dto.GetOtpRequest true "GetOtpRequest"
// @Success 	201 {object} helper.BaseHttpResponse "Success"
// @Failure 	400 {object} helper.BaseHttpResponse "Failed"
// @Failure 	409 {object} helper.BaseHttpResponse "Failed"
// @Router 		/v1/users/send-otp [post]
func (h *UserHandler) SendOtp(c *gin.Context) {
	req := new(dto.GetOtpRequest)
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			helper.GenerateBaseResponseWithValidationError(nil, false, -1, err))
		return
	}
	err = h.service.SendOtp(c, req)
	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithValidationError(nil, false, -1, err))
		return
	}
	// Call	internal SMS service
	c.JSON(http.StatusCreated, helper.GenerateBaseResponse(nil, true, 0))
}
