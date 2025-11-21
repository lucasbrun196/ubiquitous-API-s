package usecase

import (
	"go_api/model"
	"go_api/repository"
)

type GetUserUsecase struct {
	repository repository.GetUserRepository
}

func NewGetUserUsecase(repository repository.GetUserRepository) GetUserUsecase {
	return GetUserUsecase{
		repository: repository,
	}
}

func (g *GetUserUsecase) Call(id *int) model.Users {
	return g.repository.Call(id)
}
