package repository

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/lfcamarati/duo-core/domain/client/entity"
)

func NewClientRepository(tx *sql.Tx) entity.ClientRepository {
	return ClientMysqlRepository{tx}
}

type ClientMysqlRepository struct {
	Tx *sql.Tx
}

func (repository ClientMysqlRepository) GetAll() ([]entity.Client, error) {
	rows, err := repository.Tx.Query(`
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
