package repository

import (
	"go_api/initializers"
	"go_api/model"
)

type PostUserRepository struct {
}

func NewPostUsersRepository() PostUserRepository {
	return PostUserRepository{}
}

func (p *PostUserRepository) Call(params *model.Users) {
	initializers.DB.Create(params)
}
