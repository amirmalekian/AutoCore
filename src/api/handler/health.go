package handler

import (
	"net/http"
	"time"

	"github.com/amirmalekian/Mashinchi/src/api/helper"
	"github.com/gin-gonic/gin"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

// HealthCheck 	godoc
// @Summary 	Health Check
// @Description Health Check
// @Tags 		health
// @Accept  	json
// @Produce  	json
// @Success 	200 {object} helper.BaseHttpResponse "Success"
// @Failure 	400 {object} helper.BaseHttpResponse "Failed"
// @Router 		/v1/health/ [get]
func (h *Handler) Health(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, helper.GenerateBaseResponse("working!", true, 0))
}

func (h *Handler) HealthByID(ctx *gin.Context) {
	println(ctx.Params.ByName("id"))
	ctx.JSON(http.StatusOK, gin.H{"status": "ok", "caller": "get_health_byID", "time": time.Now().
		Format(time.RFC3339)})
}
