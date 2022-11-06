package usecase

import (
	"errors"

	"github.com/lfcamarati/duo-core/domain/user/infra/repository"
	"github.com/lfcamarati/duo-core/infra/security"
)

func NewLoginUsecase(
	repositoryFactory repository.UserRepositoryFactory,
	passwordEncrypt security.PasswordEncrypt,
) LoginUserUsecase {
	return LoginUserUsecase{repositoryFactory, passwordEncrypt}
}

type LoginUserUsecaseInput struct {
	Username string
	Password string
}

type LoginUserUsecaseOutput struct {
	Token *string `json:"token"`
}

type LoginUserUsecase struct {
	NewRepository   repository.UserRepositoryFactory
	PasswordEncrypt security.PasswordEncrypt
}

func (uc *LoginUserUsecase) Execute(input *LoginUserUsecaseInput) (*LoginUserUsecaseOutput, error) {
	repo := uc.NewRepository()

	user, err := repo.FindByUsername(input.Username)

	if (user == nil || err != nil) || !uc.PasswordEncrypt.CheckEncrypt(input.Password, user.Password) {
		return nil, errors.New("not authorized")
	}

	token, err := security.GenerateJWT(input.Username)

	if err != nil {
		return nil, err
	}

	return &LoginUserUsecaseOutput{token}, nil
}
