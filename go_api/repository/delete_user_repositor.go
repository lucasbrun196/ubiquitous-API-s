package repository

import (
	"go_api/initializers"
	"go_api/model"
)

type DeleteUserRepository struct {
}

func NewDeleteUserRepository() DeleteUserRepository {
	return DeleteUserRepository{}
}

func (d *DeleteUserRepository) Call(id *int) {
	initializers.DB.Delete(&model.Users{}, *id)
}
