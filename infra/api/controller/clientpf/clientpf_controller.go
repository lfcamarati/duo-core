package client

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lfcamarati/duo-core/domain/clientpf/infra/repository"
	"github.com/lfcamarati/duo-core/infra/api/handler"
	"github.com/lfcamarati/duo-core/infra/database"
	usecase "github.com/lfcamarati/duo-core/usecase/clientpf"
)

func Create(ctx *gin.Context) handler.ResponseError {
	input := new(usecase.CreateClientPfUsecaseInput)
	err := ctx.Bind(input)

	if err != nil {
		return handler.NewBadRequest("Erro ao ler dados de entrada: " + err.Error())
	}

	repository := repository.NewClientPfRepositoryFactory(database.Db)
	uc := usecase.NewCreateClientPfUseCase(repository)
	output, err := uc.Execute(input)

	if err != nil {
		return handler.NewUsecaseError("Erro ao registrar cliente: " + err.Error())
	}

	ctx.JSON(http.StatusOK, output)
	return nil
}

func Update(ctx *gin.Context) handler.ResponseError {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		return handler.NewBadRequest("Erro ao ler dados de entrada: " + err.Error())
	}

	input := new(usecase.UpdateClientPfUsecaseInput)
	err = ctx.Bind(input)

	if err != nil {
		return handler.NewBadRequest("Erro ao ler dados de entrada: " + err.Error())
	}

	input.ID = id
	repository := repository.NewClientPfRepositoryFactory(database.Db)
	uc := usecase.NewUpdateClientPfUseCase(repository)
	output, err := uc.Execute(input)

	if err != nil {
		return handler.NewUsecaseError("Erro ao atualizar cliente: " + err.Error())
	}

	ctx.JSON(http.StatusOK, output)
	return nil
}

func GetById(ctx *gin.Context) handler.ResponseError {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		return handler.NewInternalServerError("Erro ao recuperar cliente: " + err.Error())
	}

	clientPfRepo := repository.NewClientPfRepositoryFactory(database.Db)
	input := usecase.GetClientByIdUseCaseInput{ID: id}
	uc := usecase.NewGetClientByIdUseCase(clientPfRepo)
	output, err := uc.Execute(input)

	if err != nil {
		return handler.NewUsecaseError("Erro ao recuperar cliente pelo ID: " + err.Error())
	}

	ctx.JSON(http.StatusOK, output)
	return nil
}

func GetAll(ctx *gin.Context) handler.ResponseError {
	clientPfRepo := repository.NewClientPfRepositoryFactory(database.Db)
	input := usecase.GetAllClientsPfUseCaseInput{}
	uc := usecase.NewGetAllClientsPfUseCase(clientPfRepo)
	output, err := uc.Execute(input)

	if err != nil {
		return handler.NewUsecaseError("Erro ao recuperar clientes: " + err.Error())
	}

	ctx.JSON(http.StatusOK, output)
	return nil
}

func Delete(ctx *gin.Context) handler.ResponseError {
	textId := ctx.Params.ByName("id")

	if textId == "" || textId == ":id" {
		return handler.NewNotFoundError()
	}

	id, err := strconv.ParseInt(textId, 10, 64)

	if err != nil {
		return handler.NewInternalServerError("Erro ao remover cliente: " + err.Error())
	}

	clientPfRepo := repository.NewClientPfRepositoryFactory(database.Db)
	input := usecase.DeleteClientPfInput{ID: id}
	uc := usecase.NewDeleteClientPfUseCase(clientPfRepo)
	_, err = uc.Execute(input)

	if err != nil {
		return handler.NewUsecaseError("Erro ao remover cliente: " + err.Error())
	}

	ctx.JSON(http.StatusNoContent, nil)
	return nil
}
