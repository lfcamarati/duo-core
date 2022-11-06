package usecase

import (
	"context"

	"github.com/lfcamarati/duo-core/domain/user/entity"
	"github.com/lfcamarati/duo-core/domain/user/infra/repository"
)

func NewCreateUserUsecase(repositoryFactory repository.UserRepositoryFactory) CreateUserUsecase {
	return CreateUserUsecase{repositoryFactory}
}

type CreateUserUsecaseInput struct {
	Name     string
	Username string
	Password string
}

type CreateUserUsecaseOutput struct {
	ID *int64 `json:"id"`
}

type CreateUserUsecase struct {
	NewRepository repository.UserRepositoryFactory
}

func (uc *CreateUserUsecase) Execute(input *CreateUserUsecaseInput) (*CreateUserUsecaseOutput, error) {
	repo := uc.NewRepository()

	if err := repo.Begin(); err != nil {
		return nil, err
	}
	defer repo.Rollback()

	newUser := entity.NewUser(input.Name, input.Username, input.Password)
	id, err := repo.Save(context.TODO(), newUser)

	if err != nil {
		return nil, err
	}

	if err := repo.Commit(); err != nil {
		return nil, err
	}

	return &CreateUserUsecaseOutput{id}, nil
}
