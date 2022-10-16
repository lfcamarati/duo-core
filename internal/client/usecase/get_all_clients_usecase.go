package usecase

import (
	"github.com/lfcamarati/duo-core/internal/client/domain"
)

func NewGetAllClientsUseCase(repository domain.ClientRepository) *GetAllClientsUseCase {
	return &GetAllClientsUseCase{repository}
}

type GetAllClientsInput struct{}

type GetAllClientsOutput struct {
	Data *[]domain.ClientSearch `json:"data"`
}

type GetAllClientsUseCase struct {
	Repository domain.ClientRepository
}

func (uc *GetAllClientsUseCase) Execute(input GetAllClientsInput) (*GetAllClientsOutput, error) {
	clients, err := uc.Repository.GetAll()

	if err != nil {
		return nil, err
	}

	return &GetAllClientsOutput{&clients}, nil
}
