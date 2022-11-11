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

	if err := repository.Begin(); err != nil {
		return nil, err
	}
	defer repository.Rollback()

	err := repository.Delete(input.ID)

	if err != nil {
		return nil, err
	}

	if err := repository.Commit(); err != nil {
		return nil, err
	}

	return &DeleteClientPjOutput{}, nil
}
