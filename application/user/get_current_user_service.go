package user

import (
	userInfra "github.com/lfcamarati/duo-core/infra/domain/user"
)

func NewGetCurrentUserUsecase(
	factory userInfra.UserRepositoryFactory,
) *GetCurrentUserUsecase {
	return &GetCurrentUserUsecase{factory}
}

type GetCurrentUserUsecaseInput struct {
	Username string
}

type GetCurrentUserUsecaseOutput struct {
	Name string `json:"name"`
}

type GetCurrentUserUsecase struct {
	NewRepository userInfra.UserRepositoryFactory
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
