package socialmediamanagement

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/lfcamarati/duo-core/domain/socialmediamanagement"
	"github.com/lfcamarati/duo-core/infra/database"
)

func NewSocialMediaManagementRepositoryFactory(db *sql.DB) SocialMediaManagementRepositoryFactory {
	return func() socialmediamanagement.SocialMediaManagementRepository {
		return &SocialMediaManagementMysqlRepository{
			&database.GenericTransactor{Db: db},
		}
	}
}

type SocialMediaManagementRepositoryFactory func() socialmediamanagement.SocialMediaManagementRepository

type SocialMediaManagementMysqlRepository struct {
	*database.GenericTransactor
}

func (repository SocialMediaManagementMysqlRepository) Save(service socialmediamanagement.SocialMediaManagement) (*int64, error) {
	stmtService, err := repository.Tx.Prepare("INSERT INTO service (name, type, description, price) VALUES (?, ?, ?, ?)")

	if err != nil {
		return nil, err
	}

	rs, err := stmtService.Exec(service.Name, service.Type, service.Description, service.Price)

	if err != nil {
		return nil, err
	}

	id, err := rs.LastInsertId()

	if err != nil {
		return nil, err
	}

	stmt, err := repository.Tx.Prepare("INSERT INTO social_media_management (id, week_frequency, plan_type) VALUES (?, ?, ?)")

	if err != nil {
		return nil, err
	}

	_, err = stmt.Exec(service.ID, service.WeekFrequency, service.PlanType)

	if err != nil {
		return nil, err
	}

	return &id, nil
}

func (repository SocialMediaManagementMysqlRepository) Update(service socialmediamanagement.SocialMediaManagement) error {
	stmtService, err := repository.Tx.Prepare(`
		UPDATE service
		SET name = ?, type = ?, description = ?, price = ?
		WHERE id = ?
	`)

	if err != nil {
		return err
	}

	_, err = stmtService.Exec(service.Name, service.Type, service.Description, service.Price, service.ID)

	if err != nil {
		return err
	}

	stmt, err := repository.Tx.Prepare(`
		UPDATE social_media_management
		SET week_frequency = ?, plan_type = ?
		WHERE id = ?
	`)

	if err != nil {
		return err
	}

	_, err = stmt.Exec(service.WeekFrequency, service.PlanType, service.ID)

	if err != nil {
		return err
	}

	return nil
}

func (repository SocialMediaManagementMysqlRepository) GetAll() ([]socialmediamanagement.SocialMediaManagement, error) {
	rows, err := repository.Db.Query(`
		SELECT
			s.id AS "id",
			s.name AS "name",
			s.type AS "type",
			s.description AS "description",
			s.price AS "price",
			smm.week_frequency AS "weekFrequency",
			smm.plan_type AS "planType"
		FROM 
			social_media_management smm
			INNER JOIN service s ON s.id = smm.id
	`)

	if err != nil {
		return nil, err
	}

	services := make([]socialmediamanagement.SocialMediaManagement, 0)

	for rows.Next() {
		var service socialmediamanagement.SocialMediaManagement
		err := rows.Scan(&service.ID, &service.Name, &service.Type, &service.Description, &service.Price, &service.WeekFrequency, &service.PlanType)

		if err != nil {
			return nil, err
		}

		services = append(services, service)
	}

	return services, nil
}

func (repository SocialMediaManagementMysqlRepository) GetById(id int64) (*socialmediamanagement.SocialMediaManagement, error) {
	service := new(socialmediamanagement.SocialMediaManagement)

	err := repository.Db.QueryRow(`
		SELECT
			s.id AS "id",
			s.name AS "name",
			s.type AS "type",
			s.description AS "description",
			s.price AS "price",
			smm.week_frequency AS "weekFrequency",
			smm.plan_type AS "planType"
		FROM 
			social_media_management smm
			INNER JOIN service s ON s.id = smm.id
		WHERE
			smm.id = ?
	`, id).Scan(&service.ID, &service.Name, &service.Type, &service.Description, &service.Price, &service.WeekFrequency, &service.PlanType)

	if err != nil {
		return nil, err
	}

	return service, nil
}

func (repository SocialMediaManagementMysqlRepository) Delete(id int64) error {
	_, err := repository.Tx.Exec("DELETE FROM social_media_management WHERE id = ?", id)

	if err != nil {
		return err
	}

	_, err = repository.Tx.Exec("DELETE FROM service WHERE id = ?", id)

	if err != nil {
		return err
	}

	return nil
}
