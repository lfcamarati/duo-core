package client

import (
	"github.com/lfcamarati/duo-core/domain/client"
	clientInfra "github.com/lfcamarati/duo-core/infra/domain/client"
)

func NewGetAllClientsService(factory clientInfra.ClientRepositoryFactory) *GetAllClientsService {
	return &GetAllClientsService{factory}
}

type GetAllClientsService struct {
	NewRepository clientInfra.ClientRepositoryFactory
}

func (service *GetAllClientsService) Execute() ([]client.Client, error) {
	repository := service.NewRepository()
	return repository.GetAll()
}
