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

func (repository ClientMysqlRepository) createClient(client domain.Client) (*int64, error) {
	clientStmt, err := repository.Tx.Prepare("INSERT INTO client (address, email, phone) VALUES (?, ?, ?)")

	if err != nil {
		return nil, err
	}

	rs, err := clientStmt.Exec(client.Address, client.Email, client.Phone)

	if err != nil {
		return nil, err
	}

	id, err := rs.LastInsertId()

	if err != nil {
		return nil, err
	}

	return &id, nil
}
