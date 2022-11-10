package usecase

import (
	"github.com/lfcamarati/duo-core/domain/user/infra/repository"
)

func NewGetCurrentUserUsecase(
	repositoryFactory repository.UserRepositoryFactory,
) *GetCurrentUserUsecase {
	return &GetCurrentUserUsecase{repositoryFactory}
}

type GetCurrentUserUsecaseInput struct {
	Username string
}

type GetCurrentUserUsecaseOutput struct {
	Name string `json:"name"`
}

type GetCurrentUserUsecase struct {
	NewRepository repository.UserRepositoryFactory
}

func (uc *GetCurrentUserUsecase) Execute(input GetCurrentUserUsecaseInput) (*GetCurrentUserUsecaseOutput, error) {
	repo := uc.NewRepository()
	user, err := repo.FindByUsername(input.Username)

	if err != nil {
		return nil, err
	}

	return &GetCurrentUserUsecaseOutput{
		Name: user.Name,
	}, nil
}
