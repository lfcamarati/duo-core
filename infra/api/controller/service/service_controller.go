package service

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lfcamarati/duo-core/application/service"
	"github.com/lfcamarati/duo-core/infra/api/handler"
	"github.com/lfcamarati/duo-core/infra/database"
	serviceInfra "github.com/lfcamarati/duo-core/infra/domain/service"
)

func Create(ctx *gin.Context) handler.ResponseError {
	command := new(service.CreateServiceCommand)
	err := ctx.Bind(command)

	if err != nil {
		return handler.NewBadRequest("Erro ao ler dados de entrada: " + err.Error())
	}

	repository := serviceInfra.NewServiceRepositoryFactory(database.Db)
	service := service.NewCreateServiceService(repository)
	id, err := service.Execute(command)

	if err != nil {
		return handler.NewUsecaseError("erro ao cadastrar serviço: " + err.Error())
	}

	ctx.JSON(http.StatusOK, id)
	return nil
}

func Update(ctx *gin.Context) handler.ResponseError {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		return handler.NewBadRequest("Erro ao ler dados de entrada: " + err.Error())
	}

	command := new(service.UpdateServiceCommand)
	err = ctx.Bind(command)

	if err != nil {
		return handler.NewBadRequest("erro ao ler dados de entrada: " + err.Error())
	}

	command.Id = id
	repository := serviceInfra.NewServiceRepositoryFactory(database.Db)
	uc := service.NewUpdateServiceUsecase(repository)
	err = uc.Execute(command)

	if err != nil {
		return handler.NewUsecaseError("erro ao atualizar serviço: " + err.Error())
	}

	ctx.JSON(http.StatusNoContent, nil)
	return nil
}

func GetById(ctx *gin.Context) handler.ResponseError {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		return handler.NewInternalServerError("erro ao recuperar serviço: " + err.Error())
	}

	clientPfRepo := serviceInfra.NewServiceRepositoryFactory(database.Db)
	input := service.GetServiceByIdUseCaseInput{ID: id}
	uc := service.NewGetServiceByIdUseCase(clientPfRepo)
	output, err := uc.Execute(input)

	if err != nil {
		return handler.NewUsecaseError("erro ao recuperar serviço pelo ID: " + err.Error())
	}

	ctx.JSON(http.StatusOK, output)
	return nil
}

func GetAll(ctx *gin.Context) handler.ResponseError {
	clientPfRepo := serviceInfra.NewServiceRepositoryFactory(database.Db)
	input := service.GetAllServicesUseCaseInput{}
	uc := service.NewGetAllServicesUseCase(clientPfRepo)
	output, err := uc.Execute(input)

	if err != nil {
		return handler.NewUsecaseError("erro ao recuperar serviços: " + err.Error())
	}

	ctx.JSON(http.StatusOK, output)
	return nil
}

func Delete(ctx *gin.Context) handler.ResponseError {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		return handler.NewInternalServerError("erro ao remover serviço: " + err.Error())
	}

	clientPfRepo := serviceInfra.NewServiceRepositoryFactory(database.Db)
	input := service.DeleteServiceInput{ID: id}
	uc := service.NewDeleteServiceUseCase(clientPfRepo)
	_, err = uc.Execute(input)

	if err != nil {
		return handler.NewUsecaseError("erro ao remover serviço: " + err.Error())
	}

	ctx.JSON(http.StatusNoContent, nil)
	return nil
}
