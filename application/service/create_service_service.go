package service

import (
	"github.com/lfcamarati/duo-core/domain/service"
	serviceInfra "github.com/lfcamarati/duo-core/infra/domain/service"
)

func NewCreateServiceService(factory serviceInfra.ServiceRepositoryFactory) CreateServiceService {
	return CreateServiceService{factory}
}

type CreateServiceCommand struct {
	Name        string
	Type        string
	Description string
	Price       float64
}

type CreateServiceService struct {
	NewRepository serviceInfra.ServiceRepositoryFactory
}

func (s *CreateServiceService) Execute(command *CreateServiceCommand) (*int64, error) {
	repository := s.NewRepository()

	if err := repository.Begin(); err != nil {
		return nil, err
	}
	defer repository.Rollback()

	serviceType, err := service.NewServiceType(command.Type)

	if err != nil {
		return nil, err
	}

	service := service.NewService(command.Name, *serviceType, command.Description, command.Price)
	id, err := repository.Save(service)

	if err != nil {
		return nil, err
	}

	if err := repository.Commit(); err != nil {
		return nil, err
	}

	return id, nil
}
