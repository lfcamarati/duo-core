package repository

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/lfcamarati/duo-core/domain/service/entity"
	"github.com/lfcamarati/duo-core/infra/database"
)

func NewServiceRepositoryFactory(db *sql.DB) ServiceRepositoryFactory {
	return func() entity.ServiceRepository {
		return &ServiceMysqlRepository{
			&database.GenericTransactor{Db: db},
		}
	}
}

type ServiceRepositoryFactory func() entity.ServiceRepository

type ServiceMysqlRepository struct {
	*database.GenericTransactor
}

func (repository ServiceMysqlRepository) Save(service entity.Service) (*int64, error) {
	stmt, err := repository.Tx.Prepare("INSERT INTO service (title, description, price) VALUES (?, ?, ?)")

	if err != nil {
		return nil, err
	}

	rs, err := stmt.Exec(service.Title, service.Description, service.Price)

	if err != nil {
		return nil, err
	}

	id, err := rs.LastInsertId()

	if err != nil {
		return nil, err
	}

	return &id, nil
}

func (repository ServiceMysqlRepository) Update(service entity.Service) error {
	stmt, err := repository.Tx.Prepare(`
		UPDATE service
		SET title = ?, description = ?, price = ?
		WHERE id = ?
	`)

	if err != nil {
		return err
	}

	_, err = stmt.Exec(service.Title, service.Description, service.Price, service.ID)

	if err != nil {
		return err
	}

	return nil
}

func (repository ServiceMysqlRepository) GetAll() ([]entity.Service, error) {
	rows, err := repository.Db.Query(`
		SELECT
			s.id as "id",
			s.title as "title",
			s.description as "description",
			s.price as "price"
		FROM 
			service s
	`)

	if err != nil {
		return nil, err
	}

	services := make([]entity.Service, 0)

	for rows.Next() {
		var service entity.Service
		err := rows.Scan(&service.ID, &service.Title, &service.Description, &service.Price)

		if err != nil {
			return nil, err
		}

		services = append(services, service)
	}

	return services, nil
}

func (repository ServiceMysqlRepository) GetById(id int64) (*entity.Service, error) {
	service := new(entity.Service)

	err := repository.Db.QueryRow(`
		SELECT
			s.id as "id",
			s.title as "title",
			s.description as "description",
			s.price as "price"
		FROM 
			service s
		WHERE
			s.id = ?
	`, id).Scan(&service.ID, &service.Title, &service.Description, &service.Price)

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
