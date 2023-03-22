package factory

import (
	service "github.com/megalypse/golang-verifymy-backend-test/internal/data/services/user"
	usecases "github.com/megalypse/golang-verifymy-backend-test/internal/domain/usecases/user"
	repositoryFactory "github.com/megalypse/golang-verifymy-backend-test/internal/factory/repository/mysql"
	securityFactory "github.com/megalypse/golang-verifymy-backend-test/internal/factory/service"
)

var createUserUsecase usecases.CreateUser
var updateUserUsecase usecases.UpdateUser
var deleteUserUsecase usecases.DeleteUser
var findUserByIdUsecase usecases.FindUserById

func init() {
	userService := service.NewUserService(
		repositoryFactory.GetUserRepository(),
		repositoryFactory.GetMySqlUserPasswordRepository(),
		repositoryFactory.GetMySqlAddressRepository(),
		securityFactory.GetBCryptSecurityService(),
	)

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
