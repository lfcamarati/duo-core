package repository

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/lfcamarati/duo-core/domain/clientpj/entity"
	"github.com/lfcamarati/duo-core/infra/database"
)

func NewClientPjRepositoryFactory(db *sql.DB) ClientPjRepositoryFactory {
	return func() entity.ClientPjRepository {
		return &ClientPjMysqlRepository{
			&database.GenericTransactor{Db: db},
		}
	}
}

type ClientPjRepositoryFactory func() entity.ClientPjRepository

type ClientPjMysqlRepository struct {
	*database.GenericTransactor
}

func (repository ClientPjMysqlRepository) Save(clientpj entity.ClientPj) (*int64, error) {
	stmt, err := repository.Tx.Prepare("INSERT INTO client (address, email, phone, type) VALUES (?, ?, ?, ?)")

	if err != nil {
		return nil, err
	}

	rs, err := stmt.Exec(clientpj.Address, clientpj.Email, clientpj.Phone, clientpj.Type)

	if err != nil {
		return nil, err
	}

	id, err := rs.LastInsertId()

	if err != nil {
		return nil, err
	}

	clientPfStmt, err := repository.Tx.Prepare("INSERT INTO client_pj (id, corporate_name, cnpj) VALUES (?, ?, ?)")

	if err != nil {
		return nil, err
	}

	_, err = clientPfStmt.Exec(id, clientpj.CorporateName, clientpj.Cnpj)

	if err != nil {
		return nil, err
	}

	return &id, nil
}

func (repository ClientPjMysqlRepository) Update(clientpj entity.ClientPj) error {
	stmt, err := repository.Tx.Prepare(`
		UPDATE client
		SET address = ?, email = ?, phone = ?
		WHERE id = ?
	`)

	if err != nil {
		return err
	}

	_, err = stmt.Exec(clientpj.Address, clientpj.Email, clientpj.Phone, clientpj.ID)

	if err != nil {
		return err
	}

	clientPfStmt, err := repository.Tx.Prepare(`
		UPDATE client_pj
		SET corporate_name = ?, cnpj = ?
		WHERE id = ?
	`)

	if err != nil {
		return err
	}

	_, err = clientPfStmt.Exec(clientpj.CorporateName, clientpj.Cnpj, clientpj.ID)

	if err != nil {
		return err
	}

	return nil
}

func (repository ClientPjMysqlRepository) GetAll() ([]entity.ClientPj, error) {
	rows, err := repository.Db.Query(`
		SELECT
			c.id as "id",
			c.type as "type",
			pj.corporate_name as "corporateName",
			pj.cnpj as "cnpj",
			c.address as "address",
			c.email as "email",
			c.phone as "phone"
		FROM 
			client c
			inner join client_pj pj on pj.id = c.id
	`)

	if err != nil {
		return nil, err
	}

	clients := make([]entity.ClientPj, 0)

	for rows.Next() {
		var client entity.ClientPj
		err := rows.Scan(&client.ID, &client.Type, &client.CorporateName, &client.Cnpj, &client.Address, &client.Email, &client.Phone)

		if err != nil {
			return nil, err
		}

		clients = append(clients, client)
	}

	return clients, nil
}

func (repository ClientPjMysqlRepository) GetById(id int64) (*entity.ClientPj, error) {
	client := new(entity.ClientPj)

	err := repository.Db.QueryRow(`
		SELECT
			c.id as "id",
			c.type as "type",
			pj.corporate_name as "corporateName",
			pj.cnpj as "cnpj",
			c.address as "address",
			c.email as "email",
			c.phone as "phone"
		FROM 
			client c
			inner join client_pj pj on pj.id = c.id
		WHERE
			c.id = ?
	`, id).Scan(&client.ID, &client.Type, &client.CorporateName, &client.Cnpj, &client.Address, &client.Email, &client.Phone)

	if err != nil {
		return nil, err
	}

	return client, nil
}

func (repository ClientPjMysqlRepository) Delete(id int64) error {
	_, err := repository.Tx.Exec("DELETE FROM client_pj WHERE id = ?", id)

	if err != nil {
		return err
	}

	_, err = repository.Tx.Exec("DELETE FROM client WHERE id = ?", id)

	if err != nil {
		return err
	}

	return nil
}
