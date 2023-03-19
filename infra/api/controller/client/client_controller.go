package client

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lfcamarati/duo-core/application/client"
	"github.com/lfcamarati/duo-core/infra/api/handler"
	"github.com/lfcamarati/duo-core/infra/database"
	clientInfra "github.com/lfcamarati/duo-core/infra/domain/client"
)

func GetAll(ctx *gin.Context) handler.ResponseError {
	clientRepo := clientInfra.NewClientRepositoryFactory(database.Db)
	getAllClientsService := client.NewGetAllClientsService(clientRepo)
	clients, err := getAllClientsService.Execute()

	if err != nil {
		return handler.NewUsecaseError("Erro ao recuperar clientes: " + err.Error())
	}

	clientListResource := NewClientListResource(clients)
	ctx.JSON(http.StatusOK, clientListResource)
	return nil
}

func GetById(ctx *gin.Context) handler.ResponseError {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		return handler.NewInternalServerError("Erro ao recuperar cliente: " + err.Error())
	}

	repoFactory := clientInfra.NewClientRepositoryFactory(database.Db)
	getClientByIdService := client.NewGetClientByIdService(repoFactory)
	client, err := getClientByIdService.Execute(id)

	if err != nil {
		return handler.NewUsecaseError("Erro ao recuperar cliente pelo ID: " + err.Error())
	}

	ctx.JSON(http.StatusOK, NewClientResource(*client))
	return nil
}

func Create(ctx *gin.Context) handler.ResponseError {
	command := new(client.CreateClientCommand)
	err := ctx.Bind(command)

	if err != nil {
		return handler.NewBadRequest("Erro ao ler dados de entrada: " + err.Error())
	}

	repository := clientInfra.NewClientRepositoryFactory(database.Db)
	service := client.NewCreateClientService(repository)
	id, err := service.Execute(command)

	if err != nil {
		return handler.NewUsecaseError("Erro ao registrar cliente: " + err.Error())
	}

	ctx.JSON(http.StatusOK, id)
	return nil
}

func Update(ctx *gin.Context) handler.ResponseError {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		return handler.NewBadRequest("Erro ao ler dados de entrada: " + err.Error())
	}

	command := new(client.UpdateClientCommand)
	err = ctx.Bind(command)

	if err != nil {
		return handler.NewBadRequest("Erro ao ler dados de entrada: " + err.Error())
	}

	command.ID = id
	repository := clientInfra.NewClientRepositoryFactory(database.Db)
	service := client.NewUpdateClientService(repository)
	err = service.Execute(command)

	if err != nil {
		return handler.NewUsecaseError("Erro ao atualizar cliente: " + err.Error())
	}

	ctx.JSON(http.StatusNoContent, nil)
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

	clientRepo := clientInfra.NewClientRepositoryFactory(database.Db)
	command := client.DeleteClientCommand{ID: id}
	uc := client.NewDeleteClientService(clientRepo)
	err = uc.Execute(command)

	if err != nil {
		return handler.NewUsecaseError("Erro ao remover cliente: " + err.Error())
	}

	ctx.JSON(http.StatusNoContent, nil)
	return nil
}
