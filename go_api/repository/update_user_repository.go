package repository

import (
	"go_api/initializers"
	"go_api/model"
)

type UpdateUserRepository struct {
}

func NewUpdateUserRepository() UpdateUserRepository {
	return UpdateUserRepository{}
}

func (u *UpdateUserRepository) Call(params *model.Users) {
	initializers.DB.Save(&params)
}
