package client

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lfcamarati/duo-core/internal/client/infra/repository"
	"github.com/lfcamarati/duo-core/internal/client/usecase"
	"github.com/lfcamarati/duo-core/pkg/database"
)

type ErrorMessage struct {
	Message string
}

func CreateClient(ctx *gin.Context) {
	var err error

	newClientInput := new(usecase.CreateClientInput)
	err = ctx.Bind(newClientInput)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorMessage{Message: "Erro ao ler dados de entrada!"})
		return
	}

	tx, err := database.Db.BeginTx(context.TODO(), nil)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorMessage{Message: "Erro ao iniciar transação: " + err.Error()})
		return
	}
	defer tx.Rollback()

	repository := repository.NewClientMysqlRepository(tx)
	uc := usecase.NewCreateClientUseCase(repository)
	output, err := uc.Execute(*newClientInput)

	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, ErrorMessage{Message: "Erro ao cadastrar cliente: " + err.Error()})
		return
	}

	if err = tx.Commit(); err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorMessage{Message: "Erro ao gravar dados: " + err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, output)
}

func GetAll(ctx *gin.Context) {
	tx, err := database.Db.BeginTx(context.TODO(), nil)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorMessage{Message: "Erro ao iniciar transação: " + err.Error()})
		return
	}
	defer tx.Rollback()

	repository := repository.NewClientMysqlRepository(tx)
	uc := usecase.NewGetAllClientsUseCase(repository)
	output, err := uc.Execute(usecase.GetAllClientsInput{})

	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, ErrorMessage{Message: "Erro ao recuperar clientes: " + err.Error()})
		return
	}

	if err = tx.Commit(); err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorMessage{Message: "Erro ao gravar dados: " + err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, output)
}

func DeleteById(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorMessage{Message: "Erro ao remover cliente: " + err.Error()})
		return
	}

	tx, err := database.Db.BeginTx(context.TODO(), nil)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorMessage{Message: "Erro ao iniciar transação: " + err.Error()})
		return
	}
	defer tx.Rollback()

	input := usecase.DeleteClientInput{ID: id}
	repository := repository.NewClientMysqlRepository(tx)
	uc := usecase.NewDeleteClientUseCase(repository)
	output, err := uc.Execute(input)

	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, ErrorMessage{Message: "Erro ao remover cliente: " + err.Error()})
		return
	}

	if err = tx.Commit(); err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorMessage{Message: "Erro ao gravar dados: " + err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, output)
}
