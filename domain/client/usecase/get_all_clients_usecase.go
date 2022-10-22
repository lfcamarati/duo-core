package usecase

import (
	"github.com/lfcamarati/duo-core/domain/client/entity"
)

func NewGetAllClientsUseCase(repository entity.ClientRepository) *GetAllClientsUseCase {
	return &GetAllClientsUseCase{repository}
}

type GetAllClientsUsecaseInput struct{}

type ClientOutput struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type GetAllClientsUsecaseOutput struct {
	Data []ClientOutput `json:"data"`
}

type GetAllClientsUseCase struct {
	Repository entity.ClientRepository
}

func (uc *GetAllClientsUseCase) Execute(input GetAllClientsUsecaseInput) (*GetAllClientsUsecaseOutput, error) {
	clients, err := uc.Repository.GetAll()

	if err != nil {
		return nil, err
	}

	clientsOutput := make([]ClientOutput, 0)

	for _, pf := range clients {
		clientOutput := ClientOutput{
			ID:   *pf.ID,
			Name: pf.Name,
			Type: pf.Type,
		}

		clientsOutput = append(clientsOutput, clientOutput)
	}

	return &GetAllClientsUsecaseOutput{clientsOutput}, nil
}
