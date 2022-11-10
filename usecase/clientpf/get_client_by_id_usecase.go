package usecase

import (
	"github.com/lfcamarati/duo-core/domain/clientpf/infra/repository"
)

func NewGetClientByIdUseCase(factory repository.ClientPfRepositoryFactory) GetClientByIdUseCase {
	return GetClientByIdUseCase{factory}
}

type GetClientByIdUseCaseInput struct {
	ID int64
}

type GetClientByIdUseCaseOutput struct {
	ID      int64  `json:"id"`
	Name    string `json:"name"`
	Cpf     string `json:"cpf"`
	Address string `json:"address"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
}

type GetClientByIdUseCase struct {
	NewRepository repository.ClientPfRepositoryFactory
}

func (uc *GetClientByIdUseCase) Execute(input GetClientByIdUseCaseInput) (*GetClientByIdUseCaseOutput, error) {
	repository := uc.NewRepository()
	clientPf, err := repository.GetById(input.ID)

	if err != nil {
		return nil, err
	}

	return &GetClientByIdUseCaseOutput{
		ID:      *clientPf.ID,
		Name:    clientPf.Name,
		Cpf:     clientPf.Cpf,
		Address: clientPf.Address,
		Email:   clientPf.Email,
		Phone:   clientPf.Phone,
	}, nil
}
