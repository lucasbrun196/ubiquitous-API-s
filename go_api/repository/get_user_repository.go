package repository

import (
	"go_api/initializers"
	"go_api/model"
)

type GetUserRepository struct {
}

func NewGetUserRepository() GetUserRepository {
	return GetUserRepository{}
}

func (g *GetUserRepository) Call(id *int) model.Users {
	var user model.Users
	initializers.DB.First(&user, "id = ?", *id)
	return user
}
