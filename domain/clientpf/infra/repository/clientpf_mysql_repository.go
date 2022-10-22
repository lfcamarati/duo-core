package repository

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	client "github.com/lfcamarati/duo-core/domain/client/entity"
	clientRepository "github.com/lfcamarati/duo-core/domain/client/infra/repository"
	"github.com/lfcamarati/duo-core/domain/clientpf/entity"
)

func NewClientPfMysqlRepository(tx *sql.Tx) entity.ClientPfRepository {
	return ClientPfMysqlRepository{tx}
}

type ClientPfMysqlRepository struct {
	Tx *sql.Tx
}

func (repository ClientPfMysqlRepository) Save(clientpf entity.ClientPf) (*int64, error) {
	client := client.Client{
		Address: clientpf.Address,
		Email:   clientpf.Email,
		Phone:   clientpf.Phone,
		Type:    clientpf.Type,
	}

	clientRepo := clientRepository.NewClientMysqlRepository(repository.Tx)
	id, err := clientRepo.Save(client)

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

	return id, nil
}

func (repository ClientPfMysqlRepository) GetAll() ([]entity.ClientPf, error) {
	rows, err := repository.Tx.Query(`
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

	err := repository.Tx.QueryRow(`
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

	clientRepo := clientRepository.NewClientMysqlRepository(repository.Tx)
	err = clientRepo.Delete(id)

	if err != nil {
		return err
	}

	return nil
}
