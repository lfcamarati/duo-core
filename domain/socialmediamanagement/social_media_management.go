package socialmediamanagement

import (
	"errors"

	"github.com/lfcamarati/duo-core/domain/service"
)

func NewSocialMediaManagement(service service.Service, weekFrenquency int8, planTyle PlanType) SocialMediaManagement {
	return SocialMediaManagement{
		ID:            service.ID,
		Type:          service.Type,
		Name:          service.Name,
		Description:   service.Description,
		Price:         service.Price,
		WeekFrequency: weekFrenquency,
		PlanType:      planTyle,
	}
}

func NewPlanType(planTypeName string) (*PlanType, error) {
	planType := PlanType(planTypeName)

	if planType != Mensal &&
		planType != Semestral &&
		planType != Anual {
		return nil, errors.New("tipo de plano inválido")
	}

	return &planType, nil
}

type PlanType string

const (
	Mensal    PlanType = "MENSAL"
	Semestral PlanType = "SEMESTRAL"
	Anual     PlanType = "ANUAL"
)

type SocialMediaManagement struct {
	ID            *int64
	Type          service.ServiceType
	Name          string
	Description   string
	Price         float64
	WeekFrequency int8
	PlanType      PlanType
}
