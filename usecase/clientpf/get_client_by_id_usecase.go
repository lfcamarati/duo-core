package usecase

import (
	clientPf "github.com/lfcamarati/duo-core/domain/clientpf/entity"
)

func NewGetClientByIdUseCase(clientPfRepository clientPf.ClientPfRepository) GetClientByIdUseCase {
	return GetClientByIdUseCase{clientPfRepository}
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
	Repository clientPf.ClientPfRepository
}

func (uc *GetClientByIdUseCase) Execute(input GetClientByIdUseCaseInput) (*GetClientByIdUseCaseOutput, error) {
	clientPf, err := uc.Repository.GetById(input.ID)

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
