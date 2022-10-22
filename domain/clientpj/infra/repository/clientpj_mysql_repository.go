package repository

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	client "github.com/lfcamarati/duo-core/domain/client/entity"
	clientRepository "github.com/lfcamarati/duo-core/domain/client/infra/repository"
	"github.com/lfcamarati/duo-core/domain/clientpj/entity"
)

func NewClientPjRepository(tx *sql.Tx) entity.ClientPjRepository {
	return ClientPjMysqlRepository{tx}
}

type ClientPjMysqlRepository struct {
	Tx *sql.Tx
}

func (repository ClientPjMysqlRepository) Save(clientpj entity.ClientPj) (*int64, error) {
	client := client.Client{
		Address: clientpj.Address,
		Email:   clientpj.Email,
		Phone:   clientpj.Phone,
		Type:    clientpj.Type,
	}

	clientRepo := clientRepository.NewClientMysqlRepository(repository.Tx)
	id, err := clientRepo.Save(client)

	if err != nil {
		return nil, err
	}

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

	return id, nil
}

func (repository ClientPjMysqlRepository) GetAll() ([]entity.ClientPj, error) {
	rows, err := repository.Tx.Query(`
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

	err := repository.Tx.QueryRow(`
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

	clientRepo := clientRepository.NewClientMysqlRepository(repository.Tx)
	err = clientRepo.Delete(id)

	if err != nil {
		return err
	}

	return nil
}
