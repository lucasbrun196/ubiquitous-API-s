package controller

import (
	"go_api/model"
	"go_api/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DeleteUserController struct {
	usecase usecase.DeleteUserUsecase
}

func NewDeleteUserController(usecase usecase.DeleteUserUsecase) DeleteUserController {
	return DeleteUserController{usecase: usecase}
}

func (d *DeleteUserController) Call(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	d.usecase.Call(&id)
	ctx.JSON(http.StatusOK, model.Response{
		StatusCode: 200,
		Message:    "User was deleted with successfully",
	})
}
