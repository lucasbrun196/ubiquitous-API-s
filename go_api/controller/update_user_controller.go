package controller

import (
	"go_api/model"
	"go_api/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type updateUserController struct {
	usecase usecase.UpdateUserUsecase
}

func NewUpdateUserController(usecase usecase.UpdateUserUsecase) updateUserController {
	return updateUserController{
		usecase: usecase,
	}
}

func (u *updateUserController) Call(ctx *gin.Context) {
	var params model.Users
	id, _ := strconv.Atoi(ctx.Param("id"))

	ctx.BindJSON(&params)
	params = model.Users{Id: id, Name: params.Name, Email: params.Email, User: params.User, Password: params.Password}
	u.usecase.Call(&params)
	ctx.JSON(http.StatusOK, model.Response{
		StatusCode: 200,
		Message:    "User was updated with successfully",
	})
}
