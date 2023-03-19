package service

import (
	"errors"
)

func NewService(name string, serviceType ServiceType, description string, price float64) Service {
	return Service{
		Name:        name,
		Type:        serviceType,
		Description: description,
		Price:       price,
	}
}

func NewServiceType(typeName string) (*ServiceType, error) {
	serviceType := ServiceType(typeName)

	if serviceType != NormalType && serviceType != SocialMediaManagementType {
		return nil, errors.New("tipo de serviço inválido")
	}

	return &serviceType, nil
}

type Service struct {
	ID          *int64
	Type        ServiceType
	Name        string
	Description string
	Price       float64
}

func (s *Service) Update(name string, description string, price float64) Service {
	return Service{
		ID:          s.ID,
		Type:        s.Type,
		Name:        name,
		Description: description,
		Price:       price,
	}
}

type ServiceType string

const (
	NormalType                ServiceType = "NORMAL"
	SocialMediaManagementType ServiceType = "SOCIAL_MEDIA_MANAGEMENT"
)
