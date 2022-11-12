package client

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lfcamarati/duo-core/domain/client/infra/repository"
	"github.com/lfcamarati/duo-core/infra/api/handler"
	"github.com/lfcamarati/duo-core/infra/database"
	usecase "github.com/lfcamarati/duo-core/usecase/client"
)

func GetAll(ctx *gin.Context) handler.ResponseError {
	clientPfRepo := repository.NewClientRepositoryFactory(database.Db)
	input := usecase.GetAllClientsUsecaseInput{}
	uc := usecase.NewGetAllClientsUseCase(clientPfRepo)
	output, err := uc.Execute(input)

	if err != nil {
		return handler.NewUsecaseError("Erro ao recuperar clientes: " + err.Error())
	}

	ctx.JSON(http.StatusOK, output)
	return nil
}

func GetById(ctx *gin.Context) handler.ResponseError {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		return handler.NewInternalServerError("Erro ao recuperar cliente: " + err.Error())
	}

	repoFactory := repository.NewClientRepositoryFactory(database.Db)
	input := usecase.GetClientByIdUseCaseInput{ID: id}
	uc := usecase.NewGetClientByIdUseCase(repoFactory)
	output, err := uc.Execute(input)

	if err != nil {
		return handler.NewUsecaseError("Erro ao recuperar cliente pelo ID: " + err.Error())
	}

	ctx.JSON(http.StatusOK, output)
	return nil
}
