package client

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lfcamarati/duo-core/domain/client/infra/repository"
	"github.com/lfcamarati/duo-core/infra/api/handler"
	"github.com/lfcamarati/duo-core/infra/database"
	usecase "github.com/lfcamarati/duo-core/usecase/client"
)

func GetAll(ctx *gin.Context) handler.ResponseError {
	tx, err := database.Db.BeginTx(context.TODO(), nil)

	if err != nil {
		return handler.NewInternalServerError("Erro ao iniciar transação: " + err.Error())
	}
	defer tx.Rollback()

	clientPfRepo := repository.NewClientRepository(tx)
	input := usecase.GetAllClientsUsecaseInput{}
	uc := usecase.NewGetAllClientsUseCase(clientPfRepo)
	output, err := uc.Execute(input)

	if err != nil {
		return handler.NewUsecaseError("Erro ao recuperar clientes: " + err.Error())
	}

	if err = tx.Commit(); err != nil {
		return handler.NewInternalServerError("Erro ao gravar dados: " + err.Error())
	}

	ctx.JSON(http.StatusOK, output)
	return nil
}
