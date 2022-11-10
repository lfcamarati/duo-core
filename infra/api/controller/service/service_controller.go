package service

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lfcamarati/duo-core/domain/service/infra/repository"
	"github.com/lfcamarati/duo-core/infra/api/handler"
	"github.com/lfcamarati/duo-core/infra/database"
	usecase "github.com/lfcamarati/duo-core/usecase/service"
)

func Create(ctx *gin.Context) handler.ResponseError {
	input := new(usecase.CreateServiceUsecaseInput)
	err := ctx.Bind(input)

	if err != nil {
		return handler.NewBadRequest("Erro ao ler dados de entrada: " + err.Error())
	}

	repository := repository.NewServiceRepositoryFactory(database.Db)
	uc := usecase.NewCreateServiceUsecase(repository)
	output, err := uc.Execute(input)

	if err != nil {
		return handler.NewUsecaseError("erro ao cadastrar serviço: " + err.Error())
	}

	ctx.JSON(http.StatusOK, output)
	return nil
}

func Update(ctx *gin.Context) handler.ResponseError {
	input := new(usecase.UpdateServiceUsecaseInput)
	err := ctx.Bind(input)

	if err != nil {
		return handler.NewBadRequest("erro ao ler dados de entrada: " + err.Error())
	}

	repository := repository.NewServiceRepositoryFactory(database.Db)
	uc := usecase.NewUpdateServiceUsecase(repository)
	output, err := uc.Execute(*input)

	if err != nil {
		return handler.NewUsecaseError("erro ao atualizar serviço: " + err.Error())
	}

	ctx.JSON(http.StatusOK, output)
	return nil
}

func GetById(ctx *gin.Context) handler.ResponseError {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		return handler.NewInternalServerError("erro ao recuperar serviço: " + err.Error())
	}

	clientPfRepo := repository.NewServiceRepositoryFactory(database.Db)
	input := usecase.GetServiceByIdUseCaseInput{ID: id}
	uc := usecase.NewGetServiceByIdUseCase(clientPfRepo)
	output, err := uc.Execute(input)

	if err != nil {
		return handler.NewUsecaseError("erro ao recuperar serviço pelo ID: " + err.Error())
	}

	ctx.JSON(http.StatusOK, output)
	return nil
}

func GetAll(ctx *gin.Context) handler.ResponseError {
	clientPfRepo := repository.NewServiceRepositoryFactory(database.Db)
	input := usecase.GetAllServicesUseCaseInput{}
	uc := usecase.NewGetAllServicesUseCase(clientPfRepo)
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

	clientPfRepo := repository.NewServiceRepositoryFactory(database.Db)
	input := usecase.DeleteServiceInput{ID: id}
	uc := usecase.NewDeleteServiceUseCase(clientPfRepo)
	_, err = uc.Execute(input)

	if err != nil {
		return handler.NewUsecaseError("erro ao remover serviço: " + err.Error())
	}

	ctx.JSON(http.StatusNoContent, nil)
	return nil
}
