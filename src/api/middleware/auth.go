package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/amirmalekian/Mashinchi/src/api/helper"
	"github.com/amirmalekian/Mashinchi/src/config"
	"github.com/amirmalekian/Mashinchi/src/constants"
	"github.com/amirmalekian/Mashinchi/src/pkg/service_errors"
	"github.com/amirmalekian/Mashinchi/src/services"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func Authentication(cfg *config.Config) gin.HandlerFunc {
	var tokenService = services.NewTokenService(cfg)

	return func(ctx *gin.Context) {
		var err error
		claimMap := map[string]interface{}{}
		auth := ctx.GetHeader(constants.AuthorizationHeaderKey)
		token := strings.Split(auth, " ")
		if auth == "" {
			err = &service_errors.ServiceError{EndUserMessage: service_errors.TokenRequired}
		} else {
			claimMap, err = tokenService.GetClaims(token[1])
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					err = &service_errors.ServiceError{EndUserMessage: service_errors.TokenExpired}
				default:
					err = &service_errors.ServiceError{EndUserMessage: service_errors.TokenInvalid}
				}
			}
		}
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, helper.GenerateBaseResponseWithError(
				nil, false, -2, err,
			))
			return
		}
		ctx.Set(constants.UserIdKey, claimMap[constants.UserIdKey])
		ctx.Set(constants.FirstNameKey, claimMap[constants.FirstNameKey])
		ctx.Set(constants.LastNameKey, claimMap[constants.LastNameKey])
		ctx.Set(constants.UserNameKey, claimMap[constants.UserNameKey])
		ctx.Set(constants.EmailKey, claimMap[constants.EmailKey])
		ctx.Set(constants.PhoneNumberKey, claimMap[constants.PhoneNumberKey])
		ctx.Set(constants.RolesKey, claimMap[constants.RolesKey])
		ctx.Set(constants.ExpireTimeKey, claimMap[constants.ExpireTimeKey])

		ctx.Next()
	}
}

func Authorization(validRoles []string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if len(ctx.Keys) == 0 {
			ctx.AbortWithStatusJSON(http.StatusForbidden, helper.GenerateBaseResponse(nil, false, -300))
			return
		}
		rolesVal := ctx.Keys[constants.RolesKey]
		fmt.Println(rolesVal)
		if rolesVal == nil {
			ctx.AbortWithStatusJSON(http.StatusForbidden, helper.GenerateBaseResponse(nil, false, -301))
			return
		}
		roles := rolesVal.([]interface{})
		val := map[string]int{}
		for _, item := range roles {
			val[item.(string)] = 0
		}

		for _, item := range validRoles {
			if _, ok := val[item]; ok {
				ctx.Next()
				return
			}
		}
		ctx.AbortWithStatusJSON(http.StatusForbidden, helper.GenerateBaseResponse(nil, false, -301))

	}
}
