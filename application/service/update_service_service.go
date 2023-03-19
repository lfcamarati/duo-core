package service

import (
	serviceInfra "github.com/lfcamarati/duo-core/infra/domain/service"
)

func NewUpdateServiceUsecase(factory serviceInfra.ServiceRepositoryFactory) *UpdateServiceUsecase {
	return &UpdateServiceUsecase{factory}
}

type UpdateServiceCommand struct {
	Id          int64
	Name        string
	Description string
	Price       float64
}

type UpdateServiceUsecase struct {
	NewRepository serviceInfra.ServiceRepositoryFactory
}

func (uc *UpdateServiceUsecase) Execute(command *UpdateServiceCommand) error {
	repository := uc.NewRepository()

	if err := repository.Begin(); err != nil {
		return err
	}
	defer repository.Rollback()

	service, err := repository.GetById(command.Id)

	if err != nil {
		return err
	}

	service.Name = command.Name
	service.Description = command.Description
	service.Price = command.Price

	err = repository.Update(*service)

	if err != nil {
		return err
	}

	if err := repository.Commit(); err != nil {
		return err
	}

	return nil
}
