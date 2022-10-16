package repository

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/lfcamarati/duo-core/internal/client/domain"
)

func NewClientMysqlRepository(tx *sql.Tx) *ClientMysqlRepository {
	return &ClientMysqlRepository{tx}
}

type ClientMysqlRepository struct {
	Tx *sql.Tx
}

func (repository ClientMysqlRepository) SavePf(client *domain.ClientPf) (*int64, error) {
	id, err := repository.createClient(client.Client)

	if err != nil {
		return nil, err
	}

	clientPfStmt, err := repository.Tx.Prepare("INSERT INTO client_pf (id, name, cpf) VALUES (?, ?, ?)")

	if err != nil {
		return nil, err
	}

	_, err = clientPfStmt.Exec(id, client.Name, client.Cpf)

	if err != nil {
		return nil, err
	}

	return id, nil
}

func (repository ClientMysqlRepository) SavePj(client *domain.ClientPj) (*int64, error) {
	id, err := repository.createClient(client.Client)

	if err != nil {
		return nil, err
	}

	clientPfStmt, err := repository.Tx.Prepare("INSERT INTO client_pj (id, corporate_name, cnpj) VALUES (?, ?, ?)")

	if err != nil {
		return nil, err
	}

	_, err = clientPfStmt.Exec(id, client.CorporateName, client.Cnpj)

	if err != nil {
		return nil, err
	}

	return id, nil
}

func (repository ClientMysqlRepository) GetAll() ([]domain.ClientSearch, error) {
	rows, err := repository.Tx.Query(`
		SELECT
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

	clients := make([]domain.ClientSearch, 0)

	for rows.Next() {
		var client domain.ClientSearch
		err := rows.Scan(&client.Name, &client.Type)

		if err != nil {
			return nil, err
		}

		clients = append(clients, client)
	}

	return clients, nil
}

func (repository ClientMysqlRepository) createClient(client domain.Client) (*int64, error) {
	clientStmt, err := repository.Tx.Prepare("INSERT INTO client (address, email, phone, type) VALUES (?, ?, ?, ?)")

	if err != nil {
		return nil, err
	}

	rs, err := clientStmt.Exec(client.Address, client.Email, client.Phone, client.Type)

	if err != nil {
		return nil, err
	}

	id, err := rs.LastInsertId()

	if err != nil {
		return nil, err
	}

	return &id, nil
}
