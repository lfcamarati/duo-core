package usecase

import (
	"github.com/lfcamarati/duo-core/domain/clientpj/infra/repository"
)

func NewDeleteClientPjUseCase(factory repository.ClientPjRepositoryFactory) *DeleteClientPjUseCase {
	return &DeleteClientPjUseCase{factory}
}

type DeleteClientPjInput struct {
	ID int64
}

type DeleteClientPjOutput struct{}

type DeleteClientPjUseCase struct {
	NewRepository repository.ClientPjRepositoryFactory
}

func (uc *DeleteClientPjUseCase) Execute(input DeleteClientPjInput) (*DeleteClientPjOutput, error) {
	repository := uc.NewRepository()
	repository.Begin()

	err := repository.Delete(input.ID)

	if err != nil {
		repository.Rollback()
		return nil, err
	}

	repository.Commit()
	return &DeleteClientPjOutput{}, nil
}
