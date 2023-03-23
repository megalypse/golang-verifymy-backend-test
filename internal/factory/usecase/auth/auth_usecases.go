package auth

import (
	"github.com/megalypse/golang-verifymy-backend-test/internal/data/services/auth"
	authUsecases "github.com/megalypse/golang-verifymy-backend-test/internal/domain/usecases/auth"
	repositoryFactory "github.com/megalypse/golang-verifymy-backend-test/internal/factory/repository/mysql"
	"github.com/megalypse/golang-verifymy-backend-test/internal/factory/service"
)

var authUserUsecase authUsecases.AuthUser

func init() {
	authUserUsecase = auth.NewUserEmailAuth(
		repositoryFactory.GetMySqlUserPasswordRepository(),
		repositoryFactory.GetUserRepository(),
		service.GetBCryptSecurityService(),
	)
}

func GetUserEmailAuth() authUsecases.AuthUser {
	return authUserUsecase
}
