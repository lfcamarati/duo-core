package client

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/lfcamarati/duo-core/domain/client"
	"github.com/lfcamarati/duo-core/infra/database"
)

func NewClientRepositoryFactory(db *sql.DB) ClientRepositoryFactory {
	return func() client.ClientRepository {
		return &ClientMysqlRepository{
			&database.GenericTransactor{Db: db},
		}
	}
}

type ClientRepositoryFactory func() client.ClientRepository

type ClientMysqlRepository struct {
	*database.GenericTransactor
}

func (repository ClientMysqlRepository) GetAll() ([]client.Client, error) {
	rows, err := repository.Db.Query(`
		SELECT
			c.id AS "id",
			c.name AS "name",
			c.cpfCnpj AS "cpfCnpj",
			c.address AS "address",
			c.email AS "email",
			c.phone AS "phone",
			c.type AS "type"
		FROM 
			client c
	`)

	if err != nil {
		return nil, err
	}

	clients := make([]client.Client, 0)

	for rows.Next() {
		var client client.Client
		err := rows.Scan(
			&client.Id,
			&client.Name,
			&client.CpfCnpj,
			&client.Address,
			&client.Email,
			&client.Phone,
			&client.Type,
		)

		if err != nil {
			return nil, err
		}

		clients = append(clients, client)
	}

	return clients, nil
}

func (repository ClientMysqlRepository) GetById(id int64) (*client.Client, error) {
	client := new(client.Client)

	err := repository.Db.QueryRow(`
		SELECT
			c.id AS "id",
			c.name AS "name",
			c.cpfCnpj AS "cpfCnpj",
			c.address AS "address",
			c.email AS "email",
			c.phone AS "phone",
			c.type AS "type"
		FROM 
			client c
		WHERE
			c.id = ?
	`, id).Scan(
		&client.Id,
		&client.Name,
		&client.CpfCnpj,
		&client.Address,
		&client.Email,
		&client.Phone,
		&client.Type,
	)

	if err != nil {
		return nil, err
	}

	return client, nil
}

func (repository ClientMysqlRepository) Save(client client.Client) (*int64, error) {
	stmt, err := repository.Tx.Prepare(`
		INSERT INTO client (
			name, cpfCnpj, address, email, phone, type
		) VALUES (
			?, ?, ?, ?, ?, ?
		)
	`)

	if err != nil {
		return nil, err
	}

	rs, err := stmt.Exec(client.Name, client.CpfCnpj, client.Address, client.Email, client.Phone, client.Type)

	if err != nil {
		return nil, err
	}

	id, err := rs.LastInsertId()

	if err != nil {
		return nil, err
	}

	return &id, nil
}

func (repository ClientMysqlRepository) Update(client client.Client) error {
	stmt, err := repository.Tx.Prepare(`
		UPDATE client
		SET name = ?, cpfCnpj = ?, address = ?, email = ?, phone = ?, type = ?
		WHERE id = ?
	`)

	if err != nil {
		return err
	}

	_, err = stmt.Exec(
		client.Name, client.CpfCnpj, client.Address, client.Email, client.Phone, client.Type, client.Id,
	)

	if err != nil {
		return err
	}

	return nil
}

func (repository ClientMysqlRepository) Delete(id int64) error {
	_, err := repository.Tx.Exec("DELETE FROM client WHERE id = ?", id)

	if err != nil {
		return err
	}

	return nil
}
