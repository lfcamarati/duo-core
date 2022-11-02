package client

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lfcamarati/duo-core/domain/clientpf/infra/repository"
	"github.com/lfcamarati/duo-core/infra/api/httpmessage"
	"github.com/lfcamarati/duo-core/infra/database"
	usecase "github.com/lfcamarati/duo-core/usecase/clientpf"
)

func Create(ctx *gin.Context) {
	input := new(usecase.CreateClientPfUsecaseInput)
	err := ctx.Bind(input)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, httpmessage.ErrorMessage{Message: "erro ao ler dados de entrada: " + err.Error()})
		return
	}

	tx, err := database.Db.BeginTx(context.TODO(), nil)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, httpmessage.ErrorMessage{Message: "erro ao iniciar transação: " + err.Error()})
		return
	}
	defer tx.Rollback()

	repository := repository.NewClientPfRepository(tx)
	uc := usecase.NewCreateClientPfUseCase(repository)
	output, err := uc.Execute(input)

	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, httpmessage.ErrorMessage{Message: "erro ao registrar cliente: " + err.Error()})
		return
	}

	if err = tx.Commit(); err != nil {
		ctx.JSON(http.StatusInternalServerError, httpmessage.ErrorMessage{Message: "erro ao gravar dados: " + err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, output)
}

func Update(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, httpmessage.ErrorMessage{Message: "erro ao ler dados de entrada: " + err.Error()})
		return
	}

	input := new(usecase.UpdateClientPfUsecaseInput)
	err = ctx.Bind(input)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, httpmessage.ErrorMessage{Message: "erro ao ler dados de entrada: " + err.Error()})
		return
	}

	tx, err := database.Db.BeginTx(context.TODO(), nil)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, httpmessage.ErrorMessage{Message: "erro ao iniciar transação: " + err.Error()})
		return
	}
	defer tx.Rollback()

	input.ID = id
	repository := repository.NewClientPfRepository(tx)
	uc := usecase.NewUpdateClientPfUseCase(repository)
	output, err := uc.Execute(input)

	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, httpmessage.ErrorMessage{Message: "erro ao atualizar cliente: " + err.Error()})
		return
	}

	if err = tx.Commit(); err != nil {
		ctx.JSON(http.StatusInternalServerError, httpmessage.ErrorMessage{Message: "erro ao gravar dados: " + err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, output)
}

func GetById(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, httpmessage.ErrorMessage{Message: "erro ao recuperar cliente: " + err.Error()})
		return
	}

	tx, err := database.Db.BeginTx(context.TODO(), nil)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, httpmessage.ErrorMessage{Message: "erro ao iniciar transação: " + err.Error()})
		return
	}
	defer tx.Rollback()

	clientPfRepo := repository.NewClientPfRepository(tx)
	input := usecase.GetClientByIdUseCaseInput{ID: id}
	uc := usecase.NewGetClientByIdUseCase(clientPfRepo)
	output, err := uc.Execute(input)

	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, httpmessage.ErrorMessage{Message: "erro ao recuperar cliente pelo ID: " + err.Error()})
		return
	}

	if err = tx.Commit(); err != nil {
		ctx.JSON(http.StatusInternalServerError, httpmessage.ErrorMessage{Message: "erro ao gravar dados: " + err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, output)
}

func GetAll(ctx *gin.Context) {
	tx, err := database.Db.BeginTx(context.TODO(), nil)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, httpmessage.ErrorMessage{Message: "erro ao iniciar transação: " + err.Error()})
		return
	}
	defer tx.Rollback()

	clientPfRepo := repository.NewClientPfRepository(tx)
	input := usecase.GetAllClientsPfUseCaseInput{}
	uc := usecase.NewGetAllClientsPfUseCase(clientPfRepo)
	output, err := uc.Execute(input)

	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, httpmessage.ErrorMessage{Message: "erro ao recuperar clientes: " + err.Error()})
		return
	}

	if err = tx.Commit(); err != nil {
		ctx.JSON(http.StatusInternalServerError, httpmessage.ErrorMessage{Message: "erro ao gravar dados: " + err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, output)
}

func Delete(ctx *gin.Context) {
	textId := ctx.Params.ByName("id")

	if textId == "" || textId == ":id" {
		ctx.JSON(http.StatusNotFound, nil)
		return
	}

	id, err := strconv.ParseInt(textId, 10, 64)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, httpmessage.ErrorMessage{Message: "erro ao remover cliente: " + err.Error()})
		return
	}

	tx, err := database.Db.BeginTx(context.TODO(), nil)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, httpmessage.ErrorMessage{Message: "erro ao iniciar transação: " + err.Error()})
		return
	}
	defer tx.Rollback()

	clientPfRepo := repository.NewClientPfRepository(tx)
	input := usecase.DeleteClientPfInput{ID: id}
	uc := usecase.NewDeleteClientPfUseCase(clientPfRepo)
	_, err = uc.Execute(input)

	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, httpmessage.ErrorMessage{Message: "Erro ao remover cliente: " + err.Error()})
		return
	}

	if err = tx.Commit(); err != nil {
		ctx.JSON(http.StatusInternalServerError, httpmessage.ErrorMessage{Message: "erro ao remover cliente: " + err.Error()})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
