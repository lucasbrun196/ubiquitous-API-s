package controller

import (
	"go_api/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type getUserController struct {
	usecase usecase.GetUserUsecase
}

func NewGetUserController(usecase usecase.GetUserUsecase) getUserController {
	return getUserController{
		usecase: usecase,
	}
}

func (g *getUserController) Call(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	res := g.usecase.Call(&id)
	ctx.JSON(http.StatusOK, res)
}
