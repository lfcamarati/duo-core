package client

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lfcamarati/duo-core/domain/user/infra/repository"
	"github.com/lfcamarati/duo-core/infra/api/handler"
	"github.com/lfcamarati/duo-core/infra/database"
	"github.com/lfcamarati/duo-core/infra/security"
	usecase "github.com/lfcamarati/duo-core/usecase/user"
)

func Login(ctx *gin.Context) handler.ResponseError {
	input := new(usecase.LoginUserUsecaseInput)
	err := ctx.Bind(input)

	if err != nil {
		return handler.NewBadRequest("Erro ao ler dados de entrada: " + err.Error())
	}

	userRepo := repository.NewUserRepositoryFactory(database.Db)
	passwordEncrypt := security.NewDefaultPasswordEncrypt()
	uc := usecase.NewLoginUsecase(userRepo, passwordEncrypt)
	output, err := uc.Execute(input)

	if err != nil {
		return handler.NewNotAuthorizedError(err.Error())
	}

	ctx.JSON(http.StatusOK, output)
	return nil
}

func ValidateLogin(ctx *gin.Context) handler.ResponseError {
	return nil
}
