package client

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lfcamarati/duo-core/domain/clientpj/infra/repository"
	"github.com/lfcamarati/duo-core/domain/clientpj/usecase"
	"github.com/lfcamarati/duo-core/infra/database"
)

func Create(ctx *gin.Context) {
	input := new(usecase.CreateClientPjInput)
	err := ctx.Bind(input)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Errorf("erro ao ler dados de entrada: %s", err.Error()))
		return
	}

	tx, err := database.Db.BeginTx(context.TODO(), nil)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, fmt.Errorf("erro ao iniciar transação: %s", err.Error()))
		return
	}
	defer tx.Rollback()

	repository := repository.NewClientPjRepository(tx)
	uc := usecase.NewCreateClientPjUsecase(repository)
	output, err := uc.Execute(input)

	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, fmt.Errorf("erro ao cadastrar cliente: %s", err.Error()))
		return
	}

	if err = tx.Commit(); err != nil {
		ctx.JSON(http.StatusInternalServerError, fmt.Errorf("erro ao gravar dados: %s", err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, output)
}

func GetById(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, fmt.Errorf("erro ao recuperar cliente pelo ID: %s", err.Error()))
		return
	}

	tx, err := database.Db.BeginTx(context.TODO(), nil)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, fmt.Errorf("erro ao iniciar transação: %s", err.Error()))
		return
	}
	defer tx.Rollback()

	clientPjRepo := repository.NewClientPjRepository(tx)
	input := usecase.GetClientPjByIdUseCaseInput{ID: id}
	uc := usecase.NewGetClientPjByIdUseCase(clientPjRepo)
	output, err := uc.Execute(input)

	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, fmt.Errorf("erro ao recuperar cliente pelo ID: %s", err.Error()))
		return
	}

	if err = tx.Commit(); err != nil {
		ctx.JSON(http.StatusInternalServerError, fmt.Errorf("erro ao gravar dados: %s", err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, output)
}

func GetAll(ctx *gin.Context) {
	tx, err := database.Db.BeginTx(context.TODO(), nil)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, fmt.Errorf("erro ao iniciar transação: %s", err.Error()))
		return
	}
	defer tx.Rollback()

	clientPjRepo := repository.NewClientPjRepository(tx)
	input := usecase.GetAllClientsPjUseCaseInput{}
	uc := usecase.NewGetAllClientsPjUseCase(clientPjRepo)
	output, err := uc.Execute(input)

	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, fmt.Errorf("erro ao recuperar clientes: %s", err.Error()))
		return
	}

	if err = tx.Commit(); err != nil {
		ctx.JSON(http.StatusInternalServerError, fmt.Errorf("erro ao gravar dados: %s", err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, output)
}

func Delete(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, fmt.Errorf("erro ao remover cliente: %s", err.Error()))
		return
	}

	tx, err := database.Db.BeginTx(context.TODO(), nil)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, fmt.Errorf("erro ao iniciar transação: %s", err.Error()))
		return
	}
	defer tx.Rollback()

	repository := repository.NewClientPjRepository(tx)
	input := usecase.DeleteClientPjInput{ID: id}
	uc := usecase.NewDeleteClientPjUseCase(repository)
	output, err := uc.Execute(input)

	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, fmt.Errorf("erro ao remover cliente: %s", err.Error()))
		return
	}

	if err = tx.Commit(); err != nil {
		ctx.JSON(http.StatusInternalServerError, fmt.Errorf("erro ao remover cliente: %s", err.Error()))
		return
	}

	ctx.JSON(http.StatusNoContent, output)
}
