package client

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/lfcamarati/duo-core/domain/user/infra/repository"
	"github.com/lfcamarati/duo-core/infra/api/handler"
	"github.com/lfcamarati/duo-core/infra/database"
	"github.com/lfcamarati/duo-core/infra/security"
	usecase "github.com/lfcamarati/duo-core/usecase/user"
)

func Create(ctx *gin.Context) handler.ResponseError {
	input := new(usecase.CreateUserUsecaseInput)
	err := ctx.Bind(input)

	if err != nil {
		return handler.NewBadRequest("Erro ao ler dados de entrada: " + err.Error())
	}

	userRepo := repository.NewUserRepositoryFactory(database.Db)
	passwordEncrypt := security.NewDefaultPasswordEncrypt()
	uc := usecase.NewCreateUserUsecase(userRepo, passwordEncrypt)
	output, err := uc.Execute(input)

	if err != nil {
		return handler.NewUsecaseError("Erro ao registrar novo usuário: " + err.Error())
	}

	ctx.JSON(http.StatusOK, output)
	return nil
}

func ValidateLogin(ctx *gin.Context) handler.ResponseError {
	return nil
}

func GetCurrent(ctx *gin.Context) handler.ResponseError {
	authorizationHeader := ctx.Request.Header["Authorization"]
	tokenString := strings.Replace(authorizationHeader[0], "Bearer ", "", 1)
	userToken, err := security.DecodeJwt(tokenString)

	if err != nil {
		return handler.NewUsecaseError("Erro ao recuperar usuaŕio logado: " + err.Error())
	}

	repoFactory := repository.NewUserRepositoryFactory(database.Db)
	input := usecase.GetCurrentUserUsecaseInput{Username: userToken.Username}
	uc := usecase.NewGetCurrentUserUsecase(repoFactory)
	output, err := uc.Execute(input)

	if err != nil {
		return handler.NewUsecaseError("Erro ao recuperar usuaŕio logado: " + err.Error())
	}

	ctx.JSON(http.StatusOK, output)
	return nil
}
