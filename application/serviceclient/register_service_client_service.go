package serviceclient

import (
	"fmt"

	"github.com/lfcamarati/duo-core/domain/client"
	"github.com/lfcamarati/duo-core/domain/serviceclient"
	serviceClientInfra "github.com/lfcamarati/duo-core/infra/domain/serviceclient"
)

func NewRegisterServiceClientService(factory serviceClientInfra.ServiceClientRepositoryFactory) *RegisterServiceClientService {
	return &RegisterServiceClientService{factory}
}

type RegisterServiceClientCommand struct {
	ClientId client.ClientId
	Description string
	Price float64
	PeriodType string
	WeekDays *[]string
	SpecificDate *string
}

type RegisterServiceClientService struct {
	NewRepository serviceClientInfra.ServiceClientRepositoryFactory
}

func (s *RegisterServiceClientService) Execute(input *RegisterServiceClientCommand) (*int64, error) {
	repository := s.NewRepository()

	if err := repository.Begin(); err != nil {
		return nil, err
	}
	defer repository.Rollback()

	serviceClient, err := serviceclient.NewServiceClient(input.ClientId, input.Description, input.Price, input.PeriodType, input.WeekDays, input.SpecificDate)

	if err != nil {
		return nil, err
	}

	id, err := repository.Register(serviceClient)

	if err != nil {
		fmt.Printf("err: %v\n", err)
		return nil, err
	}

	if err := repository.Commit(); err != nil {
		return nil, err
	}

	return id, nil
}
