package client

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lfcamarati/duo-core/domain/client/infra/repository"
	"github.com/lfcamarati/duo-core/infra/database"
	usecase "github.com/lfcamarati/duo-core/usecase/client"
)

func GetAll(ctx *gin.Context) {
	tx, err := database.Db.BeginTx(context.TODO(), nil)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, fmt.Errorf("erro ao iniciar transação: %s", err.Error()))
		return
	}
	defer tx.Rollback()

	clientPfRepo := repository.NewClientRepository(tx)
	input := usecase.GetAllClientsUsecaseInput{}
	uc := usecase.NewGetAllClientsUseCase(clientPfRepo)
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
