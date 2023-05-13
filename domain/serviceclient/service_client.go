package serviceclient

import (
	"errors"
	"strings"
	"time"

	"github.com/lfcamarati/duo-core/domain/client"
)

var (
	ErrClientIdRequired = errors.New("cliente deve ser informado")
	ErrDescriptionRequired = errors.New("descrição deve ser informada")
	ErrPriceRequired = errors.New("preço deve ser informada")
	ErrInvalidPeriodType = errors.New("tipo de período inválido")
	ErrWeekDaysRequired       = errors.New("dias da semana deve ser informado")
	ErrSpecificDateRequired = errors.New("data específica deve ser informada")
	ErrSpecificDateInvalid = errors.New("data específica inválida")
)

func NewServiceClient(clientId client.ClientId, description string, price float64, 
	periodTypeName string, weekDays *[]string, specificDateString *string) (ServiceClient, error) {
	
	periodType, err := newServicePeriodType(periodTypeName)
	
	if err != nil {
		return ServiceClient{}, err
	}

	specificDate, err := convertSpecificDate(specificDateString)

	if err != nil {
		return ServiceClient{}, err
	}

	serviceClient := ServiceClient{
		Id: nil,
		ClientId: clientId,
		Description: description,
		Price: price,
		PeriodType: periodType,
		WeekDays: convertWeekDays(weekDays),
		SpecificDate: specificDate,
	}

	err = validate(serviceClient)

	if err != nil {
		return ServiceClient{}, err
	}

	return serviceClient, nil
}

func validate(serviceClient ServiceClient) error {
	if serviceClient.ClientId == 0 {
		return ErrClientIdRequired
	}

	if serviceClient.Description == "" {
		return ErrDescriptionRequired
	}

	if serviceClient.Price == 0 {
		return ErrPriceRequired
	}

	if serviceClient.PeriodType == WEEKLY && (serviceClient.WeekDays == nil || *serviceClient.WeekDays == "") {
		return ErrWeekDaysRequired
	}

	if serviceClient.PeriodType == SPECIFIC_DATE && (serviceClient.SpecificDate == nil) {
		return ErrSpecificDateRequired
	}
	
	return nil
}

func convertWeekDays(weekDays *[]string) *string {
	if weekDays == nil || len(*weekDays) == 0 {
		return nil
	}
	
	var concatedWeekDays string = strings.Join(*weekDays, ",")
	return &concatedWeekDays
}

func convertSpecificDate(specificDateString *string) (*time.Time, error) {
	if specificDateString == nil {
		return nil, nil
	}

	specificDate, err := time.Parse("01/02/2006", *specificDateString)

	if err != nil {
		return nil, ErrSpecificDateInvalid
	}

	return &specificDate, nil
}

func newServicePeriodType(periodTypeName string) (ServicePeriodType, error) {
	periodType := ServicePeriodType(periodTypeName)

	if periodType != WEEKLY && periodType != SPECIFIC_DATE {
		return "", ErrInvalidPeriodType
	}

	return periodType, nil
}

type ServiceClientId int64
type ServicePeriodType string

const (
	WEEKLY ServicePeriodType = "WEEKLY"
	SPECIFIC_DATE ServicePeriodType = "SPECIFIC_DATE"
)

type ServiceClient struct {
	Id *ServiceClientId
	ClientId client.ClientId
	Description string
	Price float64
	PeriodType ServicePeriodType
	WeekDays *string
	SpecificDate *time.Time
}
