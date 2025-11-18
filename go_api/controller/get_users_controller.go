package getuserscontroller

import (
	"go_api/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type getUsersController struct {
	// usecase
}

func GetGetUserControllerInstance() getUsersController {
	return getUsersController{}
}

func (g *getUsersController) GetUsersController(ctx *gin.Context) {
	mock := []model.User{
		{Id: 1, Name: "lucas", Email: "lucas@gmail.com", User: "lucas", Password: "password123"},
	}
	ctx.JSON(http.StatusOK, mock)
}
