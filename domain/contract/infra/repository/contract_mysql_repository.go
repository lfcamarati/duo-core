package repository

import (
	"database/sql"

	"github.com/lfcamarati/duo-core/domain/contract/entity"
	"github.com/lfcamarati/duo-core/infra/database"
)

func NewContractRepositoryFactory(db *sql.DB) ContractRepositoryFactory {
	return func() entity.ContractRepository {
		return &ContractMysqlRepository{
			&database.GenericTransactor{Db: db},
		}
	}
}

type ContractRepositoryFactory func() entity.ContractRepository

type ContractMysqlRepository struct {
	*database.GenericTransactor
}

func (repository ContractMysqlRepository) Save(contract entity.Contract) (*int64, error) {
	return nil, nil
}

func (repository ContractMysqlRepository) Update(contract entity.Contract) error {
	return nil
}

func (repository ContractMysqlRepository) GetAll() ([]entity.Contract, error) {
	return nil, nil
}

func (repository ContractMysqlRepository) GetById(id int64) (*entity.Contract, error) {
	return nil, nil
}

func (repository ContractMysqlRepository) Delete(id int64) error {
	return nil
}
