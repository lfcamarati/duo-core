package usecase

import (
	"context"

	"github.com/lfcamarati/duo-core/domain/user/entity"
	"github.com/lfcamarati/duo-core/domain/user/infra/repository"
	"github.com/lfcamarati/duo-core/infra/security"
)

func NewCreateUserUsecase(
	factory repository.UserRepositoryFactory,
	passwordEncrypt security.PasswordEncrypt,
) CreateUserUsecase {
	return CreateUserUsecase{factory, passwordEncrypt}
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
	NewRepository   repository.UserRepositoryFactory
	PasswordEncrypt security.PasswordEncrypt
}

func (uc *CreateUserUsecase) Execute(input *CreateUserUsecaseInput) (*CreateUserUsecaseOutput, error) {
	repo := uc.NewRepository()

	if err := repo.Begin(); err != nil {
		return nil, err
	}
	defer repo.Rollback()

	encryptedPassword := uc.PasswordEncrypt.Encrypt(input.Password)
	newUser := entity.NewUser(input.Name, input.Username, string(encryptedPassword))
	id, err := repo.Save(context.TODO(), newUser)

	if err != nil {
		return nil, err
	}

	if err := repo.Commit(); err != nil {
		return nil, err
	}

	return &CreateUserUsecaseOutput{id}, nil
}
