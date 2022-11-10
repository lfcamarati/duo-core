package usecase

import (
	"github.com/lfcamarati/duo-core/domain/clientpj/infra/repository"
)

func NewGetClientPjByIdUseCase(factory repository.ClientPjRepositoryFactory) GetClientPjByIdUseCase {
	return GetClientPjByIdUseCase{factory}
}

type GetClientPjByIdUseCaseInput struct {
	ID int64
}

type GetClientPjByIdUseCaseOutput struct {
	ID            int64  `json:"id"`
	CorporateName string `json:"corporateName"`
	Cnpj          string `json:"cnpj"`
	Address       string `json:"address"`
	Email         string `json:"email"`
	Phone         string `json:"phone"`
}

type GetClientPjByIdUseCase struct {
	NewRepository repository.ClientPjRepositoryFactory
}

func (uc *GetClientPjByIdUseCase) Execute(input GetClientPjByIdUseCaseInput) (*GetClientPjByIdUseCaseOutput, error) {
	repository := uc.NewRepository()
	clientPj, err := repository.GetById(input.ID)

	if err != nil {
		return nil, err
	}

	return &GetClientPjByIdUseCaseOutput{
		ID:            *clientPj.ID,
		CorporateName: clientPj.CorporateName,
		Cnpj:          clientPj.Cnpj,
		Address:       clientPj.Address,
		Email:         clientPj.Email,
		Phone:         clientPj.Phone,
	}, nil
}
