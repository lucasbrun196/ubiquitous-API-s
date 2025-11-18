package main

import (
	getuserscontroller "go_api/controller"

	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default()

	GetUsersController := getuserscontroller.GetGetUserControllerInstance()

	server.GET("/users", GetUsersController.GetUsersController)

	server.Run(":3001")
}
