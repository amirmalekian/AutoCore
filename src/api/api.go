package api

import (
	"fmt"

	"github.com/amirmalekian/Mashinchi/src/api/middleware"
	"github.com/amirmalekian/Mashinchi/src/api/router"
	"github.com/amirmalekian/Mashinchi/src/api/validation"
	"github.com/amirmalekian/Mashinchi/src/config"
	"github.com/amirmalekian/Mashinchi/src/docs"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitServer(cfg *config.Config) {
	r := gin.New()

	RegisterValidators(cfg)

	r.Use(middleware.DefaultStructuredLogger(cfg))
	r.Use(middleware.Cors(cfg))
	r.Use(gin.Logger(), gin.Recovery() /*middlewares.TestMiddleware()*/, middleware.LimitByRequest())

	RegisterRoutes(r, cfg)
	RegisterSwagger(r, cfg)

	r.Run(fmt.Sprintf(":%s", cfg.Server.Port))
}

func RegisterRoutes(r *gin.Engine, cfg *config.Config) {
	api := r.Group("/api")

	v1 := api.Group("/v1")
	{
		health := v1.Group("/health")
		test_router := v1.Group("/test")
		users := v1.Group("/users")

		router.Health(health)
		router.TestRouter(test_router)
		router.User(users, cfg)
	}

	v2 := api.Group("/v2")
	{
		health := v2.Group("/health")
		router.Health(health)
	}
}

func RegisterValidators(cfg *config.Config) {
	val, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		val.RegisterValidation("phone", validation.IranianPhoneNumberValidator, false)
		val.RegisterValidation("password", func(fl validator.FieldLevel) bool {
			return validation.PasswordValidator(cfg, fl)
		}, true)
	}
}

func RegisterSwagger(r *gin.Engine, cfg *config.Config) {
	docs.SwaggerInfo.Title = "Mashinchi"
	docs.SwaggerInfo.Description = "A modern marketplace for buying and selling cars."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Host = fmt.Sprintf("localhost:%s", cfg.Server.Port)
	docs.SwaggerInfo.Schemes = []string{"http"}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
