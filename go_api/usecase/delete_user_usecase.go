package usecase

import "go_api/repository"

type DeleteUserUsecase struct {
	repository repository.DeleteUserRepository
}

func NewDeleteUserUsecase(repository repository.DeleteUserRepository) DeleteUserUsecase {
	return DeleteUserUsecase{
		repository: repository,
	}
}

func (d *DeleteUserUsecase) Call(id *int) {
	d.repository.Call(id)
}
