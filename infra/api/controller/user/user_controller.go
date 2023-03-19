package user

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	usecase "github.com/lfcamarati/duo-core/application/user"
	"github.com/lfcamarati/duo-core/infra/api/handler"
	"github.com/lfcamarati/duo-core/infra/database"
	userInfra "github.com/lfcamarati/duo-core/infra/domain/user"
	"github.com/lfcamarati/duo-core/infra/security"
)

func Create(ctx *gin.Context) handler.ResponseError {
	input := new(usecase.CreateUserUsecaseInput)
	err := ctx.Bind(input)

	if err != nil {
		return handler.NewBadRequest("Erro ao ler dados de entrada: " + err.Error())
	}

	userRepo := userInfra.NewUserRepositoryFactory(database.Db)
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

	repoFactory := userInfra.NewUserRepositoryFactory(database.Db)
	input := usecase.GetCurrentUserUsecaseInput{Username: userToken.Username}
	uc := usecase.NewGetCurrentUserUsecase(repoFactory)
	output, err := uc.Execute(input)

	if err != nil {
		return handler.NewUsecaseError("Erro ao recuperar usuaŕio logado: " + err.Error())
	}

	ctx.JSON(http.StatusOK, output)
	return nil
}
