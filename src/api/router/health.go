package router

import (
	"github.com/amirmalekian/Mashinchi/src/api/handler"
	"github.com/gin-gonic/gin"
)

func Health(r *gin.RouterGroup) {
	handler := handler.NewHandler()

	r.GET("/", handler.Health)
	r.GET("/:id", handler.HealthByID)
}
