package usecase

import (
	"go_api/model"
	getusersrepository "go_api/repository"
)

type GetUsersUsecase struct {
	repository getusersrepository.GetUsersRepository
}

func NewGetUsersUsecase(repository getusersrepository.GetUsersRepository) GetUsersUsecase {
	return GetUsersUsecase{
		repository: repository,
	}
}

func (g *GetUsersUsecase) Call() []model.Users {
	return g.repository.Call()
}
