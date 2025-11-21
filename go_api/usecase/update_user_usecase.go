package usecase

import (
	"go_api/model"
	"go_api/repository"
)

type UpdateUserUsecase struct {
	repository repository.UpdateUserRepository
}

func NewUpdateUserUsecase(repository repository.UpdateUserRepository) UpdateUserUsecase {
	return UpdateUserUsecase{
		repository: repository,
	}
}

func (u *UpdateUserUsecase) Call(params *model.Users) {
	u.repository.Call(params)
}
