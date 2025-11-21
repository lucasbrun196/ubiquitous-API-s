package usecase

import (
	"go_api/model"
	"go_api/repository"
)

type PostUserUsecase struct {
	repository repository.PostUserRepository
}

func NewPostUserUsecase(repository repository.PostUserRepository) PostUserUsecase {
	return PostUserUsecase{
		repository: repository,
	}
}

func (p *PostUserUsecase) Call(params *model.Users) {
	p.repository.Call(params)
}
