package usecase

import (
	clientPf "github.com/lfcamarati/duo-core/domain/clientpf/entity"
	clientPj "github.com/lfcamarati/duo-core/domain/clientpj/entity"
)

func NewGetAllClientsUseCase(clientPfRepository clientPf.ClientPfRepository, clientPjRepository clientPj.ClientPjRepository) *GetAllClientsUseCase {
	return &GetAllClientsUseCase{clientPfRepository, clientPjRepository}
}

type GetAllClientsInput struct{}

type ClientSearch struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type GetAllClientsOutput struct {
	Data []ClientSearch `json:"data"`
}

type GetAllClientsUseCase struct {
	ClientPfRepository clientPf.ClientPfRepository
	ClientPjRepository clientPj.ClientPjRepository
}

func (uc *GetAllClientsUseCase) Execute(input GetAllClientsInput) (*GetAllClientsOutput, error) {
	clientsPf, err := uc.ClientPfRepository.GetAll()

	if err != nil {
		return nil, err
	}

	clientsPj, err := uc.ClientPjRepository.GetAll()

	if err != nil {
		return nil, err
	}

	clients := make([]ClientSearch, 0)

	for _, pf := range clientsPf {
		client := ClientSearch{
			ID:   *pf.ID,
			Name: pf.Name,
			Type: pf.Type,
		}

		clients = append(clients, client)
	}

	for _, pj := range clientsPj {
		client := ClientSearch{
			ID:   *pj.ID,
			Name: pj.CorporateName,
			Type: pj.Type,
		}

		clients = append(clients, client)
	}

	return &GetAllClientsOutput{clients}, nil
}
