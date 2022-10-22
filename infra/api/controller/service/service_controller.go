package service

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lfcamarati/duo-core/domain/service/infra/repository"
	"github.com/lfcamarati/duo-core/domain/service/usecase"
	"github.com/lfcamarati/duo-core/infra/database"
)

func Create(ctx *gin.Context) {
	input := new(usecase.CreateServiceUsecaseInput)
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

	repository := repository.NewServiceRepository(tx)
	uc := usecase.NewCreateServiceUsecase(repository)
	output, err := uc.Execute(input)

	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, fmt.Errorf("erro ao cadastrar serviço: %s", err.Error()))
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
		ctx.JSON(http.StatusInternalServerError, fmt.Errorf("erro ao recuperar serviço: %s", err.Error()))
		return
	}

	tx, err := database.Db.BeginTx(context.TODO(), nil)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, fmt.Errorf("erro ao iniciar transação: %s", err.Error()))
		return
	}
	defer tx.Rollback()

	clientPfRepo := repository.NewServiceRepository(tx)
	input := usecase.GetServiceByIdUseCaseInput{ID: id}
	uc := usecase.NewGetServiceByIdUseCase(clientPfRepo)
	output, err := uc.Execute(input)

	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, fmt.Errorf("erro ao recuperar serviço pelo ID: %s", err.Error()))
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

	clientPfRepo := repository.NewServiceRepository(tx)
	input := usecase.GetAllServicesUseCaseInput{}
	uc := usecase.NewGetAllServicesUseCase(clientPfRepo)
	output, err := uc.Execute(input)

	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, fmt.Errorf("erro ao recuperar serviços: %s", err.Error()))
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
		ctx.JSON(http.StatusInternalServerError, fmt.Errorf("erro ao remover serviço: %s", err.Error()))
		return
	}

	tx, err := database.Db.BeginTx(context.TODO(), nil)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, fmt.Errorf("erro ao iniciar transação: %s", err.Error()))
		return
	}
	defer tx.Rollback()

	clientPfRepo := repository.NewServiceRepository(tx)
	input := usecase.DeleteServiceInput{ID: id}
	uc := usecase.NewDeleteServiceUseCase(clientPfRepo)
	_, err = uc.Execute(input)

	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, fmt.Errorf("erro ao remover serviço: %s", err.Error()))
		return
	}

	if err = tx.Commit(); err != nil {
		ctx.JSON(http.StatusInternalServerError, fmt.Errorf("erro ao remover serviço: %s", err.Error()))
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
