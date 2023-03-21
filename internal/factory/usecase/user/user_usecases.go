package factory

import (
	"github.com/megalypse/golang-verifymy-backend-test/internal/data/services"
	"github.com/megalypse/golang-verifymy-backend-test/internal/domain/usecases/usecases"
	userRepositoryFactory "github.com/megalypse/golang-verifymy-backend-test/internal/factory/repository/mysql"
)

var createUserUsecase usecases.CreateUser
var updateUserUsecase usecases.UpdateUser
var deleteUserUsecase usecases.DeleteUser
var findUserByIdUsecase usecases.FindUserById

func init() {
	userService := services.NewUserService(userRepositoryFactory.GetUserRepository())

	createUserUsecase = userService
	updateUserUsecase = userService
	deleteUserUsecase = userService
	findUserByIdUsecase = userService
}

func GetCreateUserUsecase() usecases.CreateUser {
	return createUserUsecase
}

func GetUpdateUserUsecase() usecases.UpdateUser {
	return updateUserUsecase
}

func GetDeleteUserUsecase() usecases.DeleteUser {
	return deleteUserUsecase
}

func GetFindUserByIdUsecase() usecases.FindUserById {
	return findUserByIdUsecase
}
