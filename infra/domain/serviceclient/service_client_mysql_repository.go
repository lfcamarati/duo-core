package serviceclient

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/lfcamarati/duo-core/domain/serviceclient"
	"github.com/lfcamarati/duo-core/infra/database"
)

func NewServiceClientRepositoryFactory(db *sql.DB) ServiceClientRepositoryFactory {
	return func() serviceclient.ServiceClientRepository {
		return &ServiceClientMysqlRepository{
			&database.GenericTransactor{Db: db},
		}
	}
}

type ServiceClientRepositoryFactory func() serviceclient.ServiceClientRepository

type ServiceClientMysqlRepository struct {
	*database.GenericTransactor
}

func (repository ServiceClientMysqlRepository) Register(serviceClient serviceclient.ServiceClient) (*int64, error) {
	stmt, err := repository.Tx.Prepare(`
		INSERT INTO service_client (
			id_client, description, price, period_type, week_days, specific_date
		) VALUES (
			?, ?, ?, ?, ?, ?
		)
	`)

	if err != nil {
		return nil, err
	}

	rs, err := stmt.Exec(serviceClient.ClientId, serviceClient.Description, serviceClient.Price, serviceClient.PeriodType, serviceClient.WeekDays, serviceClient.SpecificDate)

	if err != nil {
		return nil, err
	}

	id, err := rs.LastInsertId()

	if err != nil {
		return nil, err
	}

	return &id, nil
}
