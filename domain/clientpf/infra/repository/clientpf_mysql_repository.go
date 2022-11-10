package repository

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/lfcamarati/duo-core/domain/clientpf/entity"
	"github.com/lfcamarati/duo-core/infra/database"
)

func NewClientPfRepositoryFactory(db *sql.DB) ClientPfRepositoryFactory {
	return func() entity.ClientPfRepository {
		return &ClientPfMysqlRepository{
			&database.GenericTransactor{Db: db},
		}
	}
}

type ClientPfRepositoryFactory func() entity.ClientPfRepository

type ClientPfMysqlRepository struct {
	*database.GenericTransactor
}

func (repository ClientPfMysqlRepository) Save(clientpf entity.ClientPf) (*int64, error) {
	stmt, err := repository.Tx.Prepare("INSERT INTO client (address, email, phone, type) VALUES (?, ?, ?, ?)")

	if err != nil {
		return nil, err
	}

	rs, err := stmt.Exec(clientpf.Address, clientpf.Email, clientpf.Phone, clientpf.Type)

	if err != nil {
		return nil, err
	}

	id, err := rs.LastInsertId()

	if err != nil {
		return nil, err
	}

	clientPfStmt, err := repository.Tx.Prepare("INSERT INTO client_pf (id, name, cpf) VALUES (?, ?, ?)")

	if err != nil {
		return nil, err
	}

	_, err = clientPfStmt.Exec(id, clientpf.Name, clientpf.Cpf)

	if err != nil {
		return nil, err
	}

	return &id, nil
}

func (repository ClientPfMysqlRepository) Update(clientpf entity.ClientPf) error {
	stmt, err := repository.Tx.Prepare(`
		UPDATE client
		SET address = ?, email = ?, phone = ?
		WHERE id = ?
	`)

	if err != nil {
		return err
	}

	_, err = stmt.Exec(clientpf.Address, clientpf.Email, clientpf.Phone, clientpf.ID)

	if err != nil {
		return err
	}

	clientPfStmt, err := repository.Tx.Prepare(`
		UPDATE client_pf
		SET name = ?, cpf = ?
		WHERE id = ?
	`)

	if err != nil {
		return err
	}

	_, err = clientPfStmt.Exec(clientpf.Name, clientpf.Cpf, clientpf.ID)

	if err != nil {
		return err
	}

	return nil
}

func (repository ClientPfMysqlRepository) GetAll() ([]entity.ClientPf, error) {
	rows, err := repository.Db.Query(`
		SELECT
			c.id as "id",
			c.type as "type",
			pf.name as "name",
			pf.cpf as "cpf",
			c.address as "address",
			c.email as "email",
			c.phone as "phone"
		FROM 
			client c
			inner join client_pf pf on pf.id = c.id
	`)

	if err != nil {
		return nil, err
	}

	clients := make([]entity.ClientPf, 0)

	for rows.Next() {
		var client entity.ClientPf
		err := rows.Scan(&client.ID, &client.Type, &client.Name, &client.Cpf, &client.Address, &client.Email, &client.Phone)

		if err != nil {
			return nil, err
		}

		clients = append(clients, client)
	}

	return clients, nil
}

func (repository ClientPfMysqlRepository) GetById(id int64) (*entity.ClientPf, error) {
	client := new(entity.ClientPf)

	err := repository.Db.QueryRow(`
		SELECT
			c.id as "id",
			c.type as "type",
			pf.name as "name",
			pf.cpf as "cpf",
			c.address as "address",
			c.email as "email",
			c.phone as "phone"
		FROM 
			client c
			inner join client_pf pf on pf.id = c.id
		WHERE
			c.id = ?
	`, id).Scan(&client.ID, &client.Type, &client.Name, &client.Cpf, &client.Address, &client.Email, &client.Phone)

	if err != nil {
		return nil, err
	}

	return client, nil
}

func (repository ClientPfMysqlRepository) Delete(id int64) error {
	_, err := repository.Tx.Exec("DELETE FROM client_pf WHERE id = ?", id)

	if err != nil {
		return err
	}

	_, err = repository.Tx.Exec("DELETE FROM client WHERE id = ?", id)

	if err != nil {
		return err
	}

	return nil
}
