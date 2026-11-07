package middleware

import (
	"net/http"

	"github.com/amirmalekian/Mashinchi/src/api/helper"
	"github.com/gin-gonic/gin"
)

func ErrorHandler(ctx *gin.Context, err any) {
	if err, ok := err.(error); ok {
		httpResponse := helper.GenerateBaseResponseWithError(nil, false, -500, err.(error))
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, httpResponse)
		return
	}
	httpResponse := helper.GenerateBaseResponseWithAnyError(nil, false, -500, err)
	ctx.AbortWithStatusJSON(http.StatusInternalServerError, httpResponse)
}
