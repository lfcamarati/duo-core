package repository

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/lfcamarati/duo-core/domain/service/entity"
)

func NewServiceRepository(tx *sql.Tx) entity.ServiceRepository {
	return ServiceMysqlRepository{tx}
}

type ServiceMysqlRepository struct {
	Tx *sql.Tx
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

func (repository ServiceMysqlRepository) GetAll() ([]entity.Service, error) {
	rows, err := repository.Tx.Query(`
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

	err := repository.Tx.QueryRow(`
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
