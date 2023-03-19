package socialmediamanagement

import "github.com/lfcamarati/duo-core/infra/database"

type SocialMediaManagementRepository interface {
	database.Transactor

	Save(smm SocialMediaManagement) (*int64, error)
	Update(smm SocialMediaManagement) error
	GetAll() ([]SocialMediaManagement, error)
	GetById(id int64) (*SocialMediaManagement, error)
	Delete(id int64) error
}
