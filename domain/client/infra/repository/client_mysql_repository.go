package repository

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/lfcamarati/duo-core/domain/client/entity"
)

func NewClientMysqlRepository(tx *sql.Tx) entity.ClientRepository {
	return ClientMysqlRepository{tx}
}

type ClientMysqlRepository struct {
	Tx *sql.Tx
}

func (repository ClientMysqlRepository) Save(client entity.Client) (*int64, error) {
	stmt, err := repository.Tx.Prepare("INSERT INTO client (address, email, phone, type) VALUES (?, ?, ?, ?)")

	if err != nil {
		return nil, err
	}

	rs, err := stmt.Exec(client.Address, client.Email, client.Phone, client.Type)

	if err != nil {
		return nil, err
	}

	id, err := rs.LastInsertId()

	if err != nil {
		return nil, err
	}

	return &id, nil
}

func (repository ClientMysqlRepository) GetAll() ([]entity.ClientSearch, error) {
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

	clients := make([]entity.ClientSearch, 0)

	for rows.Next() {
		var client entity.ClientSearch
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

	err := repository.Tx.QueryRow("SELECT c.id, c.address, c.email, c.phone, c.type FROM client c WHERE c.id = ?", id).Scan(
		&client.ID, &client.Address, &client.Email, &client.Phone, &client.Type)

	if err != nil {
		return nil, err
	}

	return client, nil
}

func (repository ClientMysqlRepository) Delete(id int64) error {
	_, err := repository.Tx.Exec("DELETE FROM client WHERE id = ?", id)

	if err != nil {
		return err
	}

	return nil
}
