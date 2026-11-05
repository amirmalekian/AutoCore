package router

import (
	"github.com/amirmalekian/Mashinchi/src/api/handler"
	"github.com/gin-gonic/gin"
)

func TestRouter(r *gin.RouterGroup) {
	h := handler.NewTestHandler()

	r.GET("/", h.Test)
	r.GET("/users", h.Users)
	r.GET("/user/:id", h.UserByID)
	r.GET("/user/:id/accounts", h.Accounts)
	r.POST("/add-user", h.AddUser)
	r.GET("/user/get-user-by-username/:username", h.UserByUsername)

	r.POST("/binder/header1", h.HeaderBinder1)
	r.POST("/binder/header2", h.HeaderBinder2)

	r.POST("/binder/query1", h.QueryBinder1)
	r.POST("/binder/query2", h.QueryBinder2)

	r.POST("/binder/uri/:id/:name", h.UriBinder)

	r.POST("/binder/body", h.BodyBinder)

	r.POST("/binder/form", h.FormBinder)

	r.POST("/binder/file", h.FileBinder)
}
