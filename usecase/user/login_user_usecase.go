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

func (i *LoginUserUsecaseInput) IsInvalid() bool {
	return i.Username == "" || i.Password == ""
}

type LoginUserUsecaseOutput struct {
	Token *string `json:"token"`
}

type LoginUserUsecase struct {
	NewRepository   repository.UserRepositoryFactory
	PasswordEncrypt security.PasswordEncrypt
}

var (
	ErrInvalidCredentials = errors.New("usuário ou senha inválidos")
	ErrFindUserByUsername = errors.New("erro ao recuperar os dados do usuário")
)

func (uc *LoginUserUsecase) Execute(input *LoginUserUsecaseInput) (*LoginUserUsecaseOutput, error) {
	if input.IsInvalid() {
		return nil, ErrInvalidCredentials
	}

	repo := uc.NewRepository()
	user, err := repo.FindByUsername(input.Username)

	if err != nil {
		return nil, ErrFindUserByUsername
	}

	if user == nil || !uc.PasswordEncrypt.CheckEncrypt(input.Password, user.Password) {
		return nil, ErrInvalidCredentials
	}

	token, err := security.GenerateJWT(input.Username)

	if err != nil {
		return nil, err
	}

	return &LoginUserUsecaseOutput{token}, nil
}
