package controller

import (
	"go_api/model"
	"go_api/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type postUserController struct {
	usecase usecase.PostUserUsecase
}

func NewPostUsersController(usecase usecase.PostUserUsecase) postUserController {
	return postUserController{
		usecase: usecase,
	}
}

func (p *postUserController) Call(ctx *gin.Context) {
	var params *model.Users
	ctx.BindJSON(&params)
	p.usecase.Call(params)
	ctx.JSON(http.StatusCreated, model.Response{
		StatusCode: 201,
		Message:    "Created",
	})
}
