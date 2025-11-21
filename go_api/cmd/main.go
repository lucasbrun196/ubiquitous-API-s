package main

import (
	controller "go_api/controller"
	"go_api/initializers"
	repository "go_api/repository"
	usecase "go_api/usecase"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnviroments()
	initializers.ConnectDb()
}

func main() {

	server := gin.Default()

	GetUsersRepository := repository.NewGetUsersRepository()
	GetUsersUsecase := usecase.NewGetUsersUsecase(GetUsersRepository)
	GetUsersController := controller.NewGetUsersController(GetUsersUsecase)

	GetUserRepository := repository.NewGetUserRepository()
	GetUserUsecase := usecase.NewGetUserUsecase(GetUserRepository)
	GetUserController := controller.NewGetUserController(GetUserUsecase)

	PostUserRepository := repository.NewPostUsersRepository()
	PostUserUsecase := usecase.NewPostUserUsecase(PostUserRepository)
	PostUserController := controller.NewPostUsersController(PostUserUsecase)

	UpdateUserRepository := repository.NewUpdateUserRepository()
	UpdateUserUsecase := usecase.NewUpdateUserUsecase(UpdateUserRepository)
	UpdateUserController := controller.NewUpdateUserController(UpdateUserUsecase)

	DeleteUserRepository := repository.NewDeleteUserRepository()
	DeleteUserUsecase := usecase.NewDeleteUserUsecase(DeleteUserRepository)
	DeleteUserController := controller.NewDeleteUserController(DeleteUserUsecase)

	server.GET("/users", GetUsersController.Call)
	server.GET("/users/:id", GetUserController.Call)
	server.POST("/users", PostUserController.Call)
	server.PUT("/users/:id", UpdateUserController.Call)
	server.DELETE("/users/:id", DeleteUserController.Call)

	server.Run(":" + initializers.Port)
}
