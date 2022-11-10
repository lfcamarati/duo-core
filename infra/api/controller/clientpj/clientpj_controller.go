package client

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lfcamarati/duo-core/domain/clientpj/infra/repository"
	"github.com/lfcamarati/duo-core/infra/api/handler"
	"github.com/lfcamarati/duo-core/infra/database"
	usecase "github.com/lfcamarati/duo-core/usecase/clientpj"
)

func Create(ctx *gin.Context) handler.ResponseError {
	input := new(usecase.CreateClientPjInput)
	err := ctx.Bind(input)

	if err != nil {
		return handler.NewBadRequest("Erro ao ler dados de entrada: " + err.Error())
	}

	repository := repository.NewClientPjRepositoryFactory(database.Db)
	uc := usecase.NewCreateClientPjUsecase(repository)
	output, err := uc.Execute(input)

	if err != nil {
		return handler.NewUsecaseError("Erro ao cadastrar cliente: " + err.Error())
	}

	ctx.JSON(http.StatusOK, output)
	return nil
}

func Update(ctx *gin.Context) handler.ResponseError {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		return handler.NewBadRequest("Erro ao ler dados de entrada: " + err.Error())
	}

	input := new(usecase.UpdateClientPjUsecaseInput)
	err = ctx.Bind(input)

	if err != nil {
		return handler.NewBadRequest("Erro ao ler dados de entrada: " + err.Error())
	}

	input.ID = id
	repository := repository.NewClientPjRepositoryFactory(database.Db)
	uc := usecase.NewUpdateClientPjUsecase(repository)
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
		return handler.NewInternalServerError("Erro ao recuperar cliente pelo ID: " + err.Error())
	}

	clientPjRepo := repository.NewClientPjRepositoryFactory(database.Db)
	input := usecase.GetClientPjByIdUseCaseInput{ID: id}
	uc := usecase.NewGetClientPjByIdUseCase(clientPjRepo)
	output, err := uc.Execute(input)

	if err != nil {
		return handler.NewUsecaseError("Erro ao recuperar cliente pelo ID: " + err.Error())
	}

	ctx.JSON(http.StatusOK, output)
	return nil
}

func GetAll(ctx *gin.Context) handler.ResponseError {
	clientPjRepo := repository.NewClientPjRepositoryFactory(database.Db)
	input := usecase.GetAllClientsPjUseCaseInput{}
	uc := usecase.NewGetAllClientsPjUseCase(clientPjRepo)
	output, err := uc.Execute(input)

	if err != nil {
		return handler.NewUsecaseError("Erro ao recuperar clientes: " + err.Error())
	}

	ctx.JSON(http.StatusOK, output)
	return nil
}

func Delete(ctx *gin.Context) handler.ResponseError {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		return handler.NewInternalServerError("Erro ao remover cliente: " + err.Error())
	}

	repository := repository.NewClientPjRepositoryFactory(database.Db)
	input := usecase.DeleteClientPjInput{ID: id}
	uc := usecase.NewDeleteClientPjUseCase(repository)
	output, err := uc.Execute(input)

	if err != nil {
		return handler.NewUsecaseError("Erro ao remover cliente: " + err.Error())
	}

	ctx.JSON(http.StatusNoContent, output)
	return nil
}
