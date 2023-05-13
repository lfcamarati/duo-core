package serviceclient

import (
	"net/http"

	"github.com/gin-gonic/gin"
	serviceClient "github.com/lfcamarati/duo-core/application/serviceclient"
	"github.com/lfcamarati/duo-core/infra/api/handler"
	"github.com/lfcamarati/duo-core/infra/database"
	serviceClientInfra "github.com/lfcamarati/duo-core/infra/domain/serviceclient"
)

func Register(ctx *gin.Context) handler.ResponseError {
	command := new(serviceClient.RegisterServiceClientCommand)
	err := ctx.Bind(command)

	if err != nil {
		return handler.NewBadRequest("Erro ao ler dados de entrada: " + err.Error())
	}

	serviceClientRepo := serviceClientInfra.NewServiceClientRepositoryFactory(database.Db)
	registerServiceClient := serviceClient.NewRegisterServiceClientService(serviceClientRepo)
	_, err = registerServiceClient.Execute(command)

	if err != nil {
		return handler.NewUsecaseError(err.Error())
	}

	ctx.JSON(http.StatusOK, nil)
	return nil
}
