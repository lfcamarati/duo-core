package usecase

import (
	"github.com/lfcamarati/duo-core/domain/clientpj/entity"
)

func NewGetClientPjByIdUseCase(clientPfRepository entity.ClientPjRepository) GetClientPjByIdUseCase {
	return GetClientPjByIdUseCase{clientPfRepository}
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
	Repository entity.ClientPjRepository
}

func (uc *GetClientPjByIdUseCase) Execute(input GetClientPjByIdUseCaseInput) (*GetClientPjByIdUseCaseOutput, error) {
	clientPj, err := uc.Repository.GetById(input.ID)

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
