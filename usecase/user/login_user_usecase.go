package usecase

import (
	"errors"

	"github.com/lfcamarati/duo-core/domain/user/infra/repository"
	"github.com/lfcamarati/duo-core/infra/security"
)

func NewLoginUsecase(repositoryFactory repository.UserRepositoryFactory) LoginUserUsecase {
	return LoginUserUsecase{repositoryFactory}
}

type LoginUserUsecaseInput struct {
	Username string
	Password string
}

type LoginUserUsecaseOutput struct {
	Token *string `json:"token"`
}

type LoginUserUsecase struct {
	NewRepository repository.UserRepositoryFactory
}

func (uc *LoginUserUsecase) Execute(input *LoginUserUsecaseInput) (*LoginUserUsecaseOutput, error) {
	if input.Username != "teste" || input.Password != "teste" {
		return nil, errors.New("not authorized")
	}

	token, err := security.GenerateJWT(input.Username)

	if err != nil {
		return nil, err
	}

	return &LoginUserUsecaseOutput{token}, nil
}
