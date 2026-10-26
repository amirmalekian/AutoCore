package api

import (
	"fmt"

	"github.com/amirmalekian/Mashinchi/src/api/middleware"
	"github.com/amirmalekian/Mashinchi/src/api/router"
	"github.com/amirmalekian/Mashinchi/src/config"
	"github.com/gin-gonic/gin"
)

func InitServer(cfg *config.Config) {
	r := gin.New()

	r.Use(middleware.Cors(cfg))
	r.Use(gin.Logger(), gin.Recovery())

	RegisterRoutes(r)

	r.Run(fmt.Sprintf(":%s", cfg.Server.Port))
}

func RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api")

	v1 := api.Group("/v1")
	{
		health := v1.Group("/health")
		test_router := v1.Group("/test")

		router.Health(health)
		router.TestRouter(test_router)
	}

	v2 := api.Group("/v2")
	{
		health := v2.Group("/health")
		router.Health(health)
	}
}
