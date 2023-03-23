package auth

import (
	signin "github.com/megalypse/golang-verifymy-backend-test/internal/data/services/signin"
	authUsecases "github.com/megalypse/golang-verifymy-backend-test/internal/domain/usecases/auth"
	repositoryFactory "github.com/megalypse/golang-verifymy-backend-test/internal/factory/repository/mysql"
	"github.com/megalypse/golang-verifymy-backend-test/internal/factory/service"
)

var authUserUsecase authUsecases.UserSignIn

func init() {
	authUserUsecase = signin.NewUserEmailAuth(
		repositoryFactory.GetMySqlUserPasswordRepository(),
		repositoryFactory.GetUserRepository(),
		service.GetBCryptSecurityService(),
	)
}

func GetUserEmailAuth() authUsecases.UserSignIn {
	return authUserUsecase
}
