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

// RegisterByUsername 	godoc
// @Summary 			RegisterByUsername
// @Description 		RegisterByUsername
// @Tags 				Users
// @Accept  			json
// @Produce  			json
// @Param 				Request body dto.RegisterUserByUsernameRequest true "RegisterUserByUsernameRequest"
// @Success 			201 {object} helper.BaseHttpResponse "Success"
// @Failure 			400 {object} helper.BaseHttpResponse "Failed"
// @Failure 			409 {object} helper.BaseHttpResponse "Failed"
// @Router 				/v1/users/register-by-username [post]
func (h *UserHandler) RegisterByUsername(ctx *gin.Context) {
	req := new(dto.RegisterUserByUsernameRequest)
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			helper.GenerateBaseResponseWithValidationError(nil, false, -1, err))
		return
	}
	err = h.service.RegisterByUsername(req)
	if err != nil {
		ctx.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithValidationError(nil, false, -1, err))
		return
	}

	ctx.JSON(http.StatusCreated, helper.GenerateBaseResponse(nil, true, 0))
}

// LoginByUsername 	godoc
// @Summary 		LoginByUsername
// @Description 	LoginByUsername
// @Tags 			Users
// @Accept  		json
// @Produce  		json
// @Param 			Request body dto.LoginByUsernameRequest true "LoginByUsernameRequest"
// @Success 		201 {object} helper.BaseHttpResponse "Success"
// @Failure 		400 {object} helper.BaseHttpResponse "Failed"
// @Failure 		409 {object} helper.BaseHttpResponse "Failed"
// @Router 			/v1/users/login-by-username [post]
func (h *UserHandler) LoginByUsername(ctx *gin.Context) {
	req := new(dto.LoginByUsernameRequest)
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			helper.GenerateBaseResponseWithValidationError(nil, false, -1, err))
		return
	}
	token, err := h.service.LoginByUsername(req)
	if err != nil {
		ctx.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithValidationError(nil, false, -1, err))
		return
	}

	ctx.JSON(http.StatusCreated, helper.GenerateBaseResponse(token, true, 0))
}

// RegisterLoginByPhoneNumber 	godoc
// @Summary 					RegisterLoginByPhoneNumber
// @Description 				RegisterLoginByPhoneNumber
// @Tags 						Users
// @Accept  					json
// @Produce  					json
// @Param 						Request body dto.RegisterLoginByPhoneRequest true "RegisterLoginByPhoneRequest"
// @Success 					201 {object} helper.BaseHttpResponse "Success"
// @Failure 					400 {object} helper.BaseHttpResponse "Failed"
// @Failure 					409 {object} helper.BaseHttpResponse "Failed"
// @Router 						/v1/users/login-by-phone [post]
func (h *UserHandler) RegisterLoginByPhoneNumber(ctx *gin.Context) {
	req := new(dto.RegisterLoginByPhoneRequest)
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			helper.GenerateBaseResponseWithValidationError(nil, false, -1, err))
		return
	}
	token, err := h.service.RegisterLoginByPhoneNumber(ctx, req)
	if err != nil {
		ctx.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithValidationError(nil, false, -1, err))
		return
	}

	ctx.JSON(http.StatusCreated, helper.GenerateBaseResponse(token, true, 0))
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
