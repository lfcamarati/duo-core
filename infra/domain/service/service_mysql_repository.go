package service

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/lfcamarati/duo-core/domain/service"
	"github.com/lfcamarati/duo-core/infra/database"
)

func NewServiceRepositoryFactory(db *sql.DB) ServiceRepositoryFactory {
	return func() service.ServiceRepository {
		return &ServiceMysqlRepository{
			&database.GenericTransactor{Db: db},
		}
	}
}

type ServiceRepositoryFactory func() service.ServiceRepository

type ServiceMysqlRepository struct {
	*database.GenericTransactor
}

func (repository ServiceMysqlRepository) Save(service service.Service) (*int64, error) {
	stmt, err := repository.Tx.Prepare("INSERT INTO service (name, type, description, price) VALUES (?, ?, ?, ?)")

	if err != nil {
		return nil, err
	}

	rs, err := stmt.Exec(service.Name, service.Type, service.Description, service.Price)

	if err != nil {
		return nil, err
	}

	id, err := rs.LastInsertId()

	if err != nil {
		return nil, err
	}

	return &id, nil
}

func (repository ServiceMysqlRepository) Update(service service.Service) error {
	stmt, err := repository.Tx.Prepare(`
		UPDATE service
		SET name = ?, type = ?, description = ?, price = ?
		WHERE id = ?
	`)

	if err != nil {
		return err
	}

	_, err = stmt.Exec(service.Name, service.Type, service.Description, service.Price, service.ID)

	if err != nil {
		return err
	}

	return nil
}

func (repository ServiceMysqlRepository) GetAll() ([]service.Service, error) {
	rows, err := repository.Db.Query(`
		SELECT
			s.id as "id",
			s.name as "name",
			s.type as "type",
			s.description as "description",
			s.price as "price"
		FROM 
			service s
	`)

	if err != nil {
		return nil, err
	}

	services := make([]service.Service, 0)

	for rows.Next() {
		var service service.Service
		err := rows.Scan(&service.ID, &service.Name, &service.Type, &service.Description, &service.Price)

		if err != nil {
			return nil, err
		}

		services = append(services, service)
	}

	return services, nil
}

func (repository ServiceMysqlRepository) GetById(id int64) (*service.Service, error) {
	service := new(service.Service)

	err := repository.Db.QueryRow(`
		SELECT
			s.id as "id",
			s.name as "name",
			s.type as "type",
			s.description as "description",
			s.price as "price"
		FROM 
			service s
		WHERE
			s.id = ?
	`, id).Scan(&service.ID, &service.Name, &service.Type, &service.Description, &service.Price)

	if err != nil {
		return nil, err
	}

	return service, nil
}

func (repository ServiceMysqlRepository) Delete(id int64) error {
	_, err := repository.Tx.Exec("DELETE FROM service WHERE id = ?", id)

	if err != nil {
		return err
	}

	return nil
}
