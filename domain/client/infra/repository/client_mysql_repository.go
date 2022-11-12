package repository

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/lfcamarati/duo-core/domain/client/entity"
	"github.com/lfcamarati/duo-core/infra/database"
)

func NewClientRepositoryFactory(db *sql.DB) ClientRepositoryFactory {
	return func() entity.ClientRepository {
		return &ClientMysqlRepository{
			&database.GenericTransactor{Db: db},
		}
	}
}

type ClientRepositoryFactory func() entity.ClientRepository

type ClientMysqlRepository struct {
	*database.GenericTransactor
}

func (repository ClientMysqlRepository) GetAll() ([]entity.Client, error) {
	rows, err := repository.Db.Query(`
		SELECT
			c.id as "id",
			CASE
				WHEN c.type = 'PF' THEN pf.name
				ELSE pj.corporate_name
			END as "name",
			c.type as "type"
		FROM 
			client c
			left join client_pf pf on pf.id = c.id
			left join client_pj pj on pj.id = c.id
	`)

	if err != nil {
		return nil, err
	}

	clients := make([]entity.Client, 0)

	for rows.Next() {
		var client entity.Client
		err := rows.Scan(&client.ID, &client.Name, &client.Type)

		if err != nil {
			return nil, err
		}

		clients = append(clients, client)
	}

	return clients, nil
}

func (repository ClientMysqlRepository) GetById(id int64) (*entity.Client, error) {
	client := new(entity.Client)

	err := repository.Db.QueryRow(`
		SELECT
			c.id as "id",
			CASE
				WHEN c.type = 'PF' THEN pf.name
				ELSE pj.corporate_name
			END as "name",
			c.type as "type"
		FROM 
			client c
			left join client_pf pf on pf.id = c.id
			left join client_pj pj on pj.id = c.id
		WHERE
			c.id = ?
	`, id).Scan(&client.ID, &client.Name, &client.Type)

	if err != nil {
		return nil, err
	}

	return client, nil
}
