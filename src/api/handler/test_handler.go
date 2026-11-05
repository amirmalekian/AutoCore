package handler

import (
	"net/http"

	"github.com/amirmalekian/Mashinchi/src/api/helper"
	"github.com/gin-gonic/gin"
)

type TestHandler struct{}

type header struct {
	UserId  string
	Browser string
}

type personData struct {
	FirstName   string `json:"first_name" binding:"required,alpha,min=4,max=10"`
	LastName    string `json:"last_name" binding:"required,alpha,min=6,max=20"`
	PhoneNumber string `json:"phone_number" binding:"required,phone,min=11,max=11"`
}

func NewTestHandler() *TestHandler {
	return &TestHandler{}
}

func (t *TestHandler) Test(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, helper.GenerateBaseResponse("Test", true, 0))
}

func (t *TestHandler) Users(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, helper.GenerateBaseResponse("Users", true, 0))
}

// UserByID 	godoc
// @Summary 	UserByID
// @Description UserByID
// @Tags 		Test
// @Accept  	json
// @Produce  	json
// @Param 		id 	path int true "user id"
// @Success 	200 {object} helper.BaseHttpResponse "Success"
// @Failure 	400 {object} helper.BaseHttpResponse "Failed"
// @Router 		/v1/test/user/{id} [get]
func (t *TestHandler) UserByID(ctx *gin.Context) {
	id := ctx.Param("id")
	ctx.JSON(http.StatusOK, helper.GenerateBaseResponse(gin.H{
		"result": "UserById",
		"id":     id,
	}, true, 0))
}

func (t *TestHandler) Accounts(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, helper.GenerateBaseResponse(gin.H{
		"result": "Accounts",
	}, true, 0))
}

func (t *TestHandler) AddUser(ctx *gin.Context) {
	id := ctx.Param("id")
	ctx.JSON(http.StatusOK, helper.GenerateBaseResponse(gin.H{
		"result": "AddUser",
		"id":     id,
	}, true, 0))
}

func (t *TestHandler) UserByUsername(ctx *gin.Context) {
	username := ctx.Param("username")
	ctx.JSON(http.StatusOK, helper.GenerateBaseResponse(gin.H{
		"result":   "UserByUsername",
		"username": username,
	}, true, 0))
}

func (t *TestHandler) HeaderBinder1(ctx *gin.Context) {
	userId := ctx.GetHeader("UserId")

	ctx.JSON(http.StatusOK, helper.GenerateBaseResponse(gin.H{
		"result": "HeaderBinder1",
		"userId": userId,
	}, true, 0))
}

func (t *TestHandler) HeaderBinder2(ctx *gin.Context) {
	header := header{}
	ctx.BindHeader(&header)

	ctx.JSON(http.StatusOK, helper.GenerateBaseResponse(gin.H{
		"result": "HeaderBinder1",
		"header": header,
	}, true, 0))
}

func (t *TestHandler) QueryBinder1(ctx *gin.Context) {
	id := ctx.Query("id")
	name := ctx.Query("name")

	ctx.JSON(http.StatusOK, helper.GenerateBaseResponse(gin.H{
		"result": "QueryBinder1",
		"id":     id,
		"name":   name,
	}, true, 0))
}

func (t *TestHandler) QueryBinder2(ctx *gin.Context) {
	ids := ctx.QueryArray("id")
	name := ctx.Query("name")

	ctx.JSON(http.StatusOK, helper.GenerateBaseResponse(gin.H{
		"result": "QueryBinder2",
		"ids":    ids,
		"name":   name,
	}, true, 0))
}

func (t *TestHandler) UriBinder(ctx *gin.Context) {
	id := ctx.Param("id")
	name := ctx.Param("name")

	ctx.JSON(http.StatusOK, helper.GenerateBaseResponse(gin.H{
		"result": "UriBinder",
		"id":     id,
		"name":   name,
	}, true, 0))
}

// BodyBinder 	godoc
// @Summary 	BodyBinder
// @Description BodyBinder
// @Tags 		Test
// @Accept  	json
// @Produce  	json
// @Param 		person 	body personData true "person data"
// @Success 	200 {object} helper.BaseHttpResponse "Success"
// @Failure 	400 {object} helper.BaseHttpResponse "Failed"
// @Router 		/v1/test/binder/body [post]
// @Security 	AuthBearer
func (t *TestHandler) BodyBinder(ctx *gin.Context) {
	p := personData{}
	err := ctx.ShouldBindJSON(&p)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithValidationError(nil,
			false, helper.ValidationError, err))
		return
	}
	ctx.JSON(http.StatusOK, helper.GenerateBaseResponse(gin.H{
		"result": "BodyBinder",
		"person": p,
	}, true, 0))
}

func (t *TestHandler) FormBinder(ctx *gin.Context) {
	p := personData{}
	ctx.ShouldBind(&p)

	ctx.JSON(http.StatusOK, helper.GenerateBaseResponse(gin.H{
		"result": "FormBinder",
		"person": p,
	}, true, 0))
}

func (t *TestHandler) FileBinder(ctx *gin.Context) {
	file, _ := ctx.FormFile("file")
	err := ctx.SaveUploadedFile(file, "file")
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError,
			helper.GenerateBaseResponseWithError(nil, false, helper.ValidationError, err))
		return
	}

	ctx.JSON(http.StatusOK, helper.GenerateBaseResponse(gin.H{
		"result": "FileBinder",
		"file":   file.Filename,
	}, true, 0))
}
