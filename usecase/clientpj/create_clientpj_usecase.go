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
	clientPj := entity.NewClientPj(input.CorporateName, input.Cnpj, input.Address, input.Email, input.Phone)

	repository := uc.NewRepository()
	repository.Begin()

	ID, err := repository.Save(clientPj)

	if err != nil {
		repository.Rollback()
		return nil, err
	}

	repository.Commit()
	return &CreateClientPjOutput{ID}, nil
}
