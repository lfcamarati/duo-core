package usecase

import (
	"github.com/lfcamarati/duo-core/domain/clientpj/entity"
	"github.com/lfcamarati/duo-core/domain/clientpj/infra/repository"
)

func NewCreateClientPjUsecase(factory repository.ClientPjRepositoryFactory) *CreateClientPjUseCase {
	return &CreateClientPjUseCase{factory}
}

type CreateClientPjInput struct {
	Type          string
	CorporateName string
	Cnpj          string
	Address       string
	Email         string
	Phone         string
}

type CreateClientPjOutput struct {
	ID *int64
}

type CreateClientPjUseCase struct {
	NewRepository repository.ClientPjRepositoryFactory
}

func (uc *CreateClientPjUseCase) Execute(input *CreateClientPjInput) (*CreateClientPjOutput, error) {
	repository := uc.NewRepository()

	if err := repository.Begin(); err != nil {
		return nil, err
	}
	defer repository.Rollback()

	clientPj := entity.NewClientPj(input.CorporateName, input.Cnpj, input.Address, input.Email, input.Phone)
	ID, err := repository.Save(clientPj)

	if err != nil {
		return nil, err
	}

	if err := repository.Commit(); err != nil {
		return nil, err
	}

	return &CreateClientPjOutput{ID}, nil
}
