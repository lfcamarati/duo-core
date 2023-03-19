package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	userService "github.com/lfcamarati/duo-core/application/user"
	"github.com/lfcamarati/duo-core/infra/api/handler"
	"github.com/lfcamarati/duo-core/infra/database"
	userInfra "github.com/lfcamarati/duo-core/infra/domain/user"
	"github.com/lfcamarati/duo-core/infra/security"
)

func Login(ctx *gin.Context) handler.ResponseError {
	input := new(userService.LoginUserUsecaseInput)
	err := ctx.Bind(input)

	if err != nil {
		return handler.NewBadRequest("Erro ao ler dados de entrada: " + err.Error())
	}

	userRepo := userInfra.NewUserRepositoryFactory(database.Db)
	passwordEncrypt := security.NewDefaultPasswordEncrypt()
	uc := userService.NewLoginUsecase(userRepo, passwordEncrypt)
	output, err := uc.Execute(input)

	if err != nil {
		if err == userService.ErrInvalidCredentials {
			return handler.NewNotAuthorizedError(err.Error())
		} else {
			return handler.NewInternalServerError(err.Error())
		}
	}

	ctx.JSON(http.StatusOK, output)
	return nil
}

func ValidateLogin(ctx *gin.Context) handler.ResponseError {
	return nil
}
