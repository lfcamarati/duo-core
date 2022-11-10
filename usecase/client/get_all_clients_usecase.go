package usecase

import (
	"github.com/lfcamarati/duo-core/domain/client/infra/repository"
)

func NewGetAllClientsUseCase(repository repository.ClientRepositoryFactory) *GetAllClientsUseCase {
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
	NewRepository repository.ClientRepositoryFactory
}

func (uc *GetAllClientsUseCase) Execute(input GetAllClientsUsecaseInput) (*GetAllClientsUsecaseOutput, error) {
	repository := uc.NewRepository()
	clients, err := repository.GetAll()

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
