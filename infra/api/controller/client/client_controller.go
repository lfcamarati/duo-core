package client

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lfcamarati/duo-core/domain/client/usecase"
	"github.com/lfcamarati/duo-core/infra/database"

	clientRepository "github.com/lfcamarati/duo-core/domain/client/infra/repository"
	clientPfRepository "github.com/lfcamarati/duo-core/domain/clientpf/infra/repository"
	clientPfUsecase "github.com/lfcamarati/duo-core/domain/clientpf/usecase"
	clientPjRepository "github.com/lfcamarati/duo-core/domain/clientpj/infra/repository"
	clientPjUsecase "github.com/lfcamarati/duo-core/domain/clientpj/usecase"
)

type CreateClientRequest struct {
	Type          string
	Name          *string
	Cpf           *string
	CorporateName *string
	Cnpj          *string
	Address       *string
	Email         *string
	Phone         *string
}

func CreateClient(ctx *gin.Context) {
	createClientRequest := new(CreateClientRequest)
	err := ctx.Bind(createClientRequest)

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

	if createClientRequest.Type == "PF" {
		repository := clientPfRepository.NewClientPfMysqlRepository(tx)
		uc := clientPfUsecase.NewCreateClientPfUseCase(repository)
		input := clientPfUsecase.CreateClientPfInput{
			Name:    *createClientRequest.Name,
			Cpf:     *createClientRequest.Cpf,
			Address: *createClientRequest.Address,
			Email:   *createClientRequest.Email,
			Phone:   *createClientRequest.Phone,
		}
		output, err := uc.Execute(input)

		if err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, fmt.Errorf("erro ao cadastrar cliente: %s", err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, output)
	} else {
		repository := clientPjRepository.NewClientPjMysqlRepository(tx)
		uc := clientPjUsecase.NewCreateClientPjUsecase(repository)
		input := clientPjUsecase.CreateClientPjInput{
			CorporateName: *createClientRequest.CorporateName,
			Cnpj:          *createClientRequest.Cnpj,
			Address:       *createClientRequest.Address,
			Email:         *createClientRequest.Email,
			Phone:         *createClientRequest.Phone,
		}
		output, err := uc.Execute(input)

		if err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, fmt.Errorf("erro ao cadastrar cliente: %s", err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, output)
	}

	if err = tx.Commit(); err != nil {
		ctx.JSON(http.StatusInternalServerError, fmt.Errorf("erro ao gravar dados: %s", err.Error()))
		return
	}
}

func GetAll(ctx *gin.Context) {
	tx, err := database.Db.BeginTx(context.TODO(), nil)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, fmt.Errorf("erro ao iniciar transação: %s", err.Error()))
		return
	}
	defer tx.Rollback()

	repository := clientRepository.NewClientMysqlRepository(tx)
	uc := usecase.NewGetAllClientsUseCase(repository)
	output, err := uc.Execute(usecase.GetAllClientsInput{})

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

func DeleteById(ctx *gin.Context) {
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

	clientRepo := clientRepository.NewClientMysqlRepository(tx)
	client, err := clientRepo.GetById(id)

	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, fmt.Errorf("erro ao remover cliente: %s", err.Error()))
		return
	}

	if client.IsPf() {
		repository := clientPfRepository.NewClientPfMysqlRepository(tx)
		input := clientPfUsecase.DeleteClientPfInput{ID: id}
		uc := clientPfUsecase.NewDeleteClientPfUseCase(repository)
		output, err := uc.Execute(input)

		if err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, fmt.Errorf("erro ao remover cliente: %s", err.Error()))
			return
		}

		ctx.JSON(http.StatusNoContent, output)
	} else {
		repository := clientPjRepository.NewClientPjMysqlRepository(tx)
		input := clientPjUsecase.DeleteClientPjInput{ID: id}
		uc := clientPjUsecase.NewDeleteClientPjUseCase(repository)
		output, err := uc.Execute(input)

		if err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, fmt.Errorf("erro ao remover cliente: %s", err.Error()))
			return
		}

		ctx.JSON(http.StatusNoContent, output)
	}

	if err = tx.Commit(); err != nil {
		ctx.JSON(http.StatusInternalServerError, fmt.Errorf("erro ao remover cliente: %s", err.Error()))
		return
	}
}
