package factory

import (
	userUsecaseFactory "github.com/megalypse/golang-verifymy-backend-test/internal/factory/usecase/user"
	"github.com/megalypse/golang-verifymy-backend-test/internal/presentation/http/controllers"
	"github.com/megalypse/golang-verifymy-backend-test/internal/presentation/http/controllers/auth"
	"github.com/megalypse/golang-verifymy-backend-test/internal/presentation/http/controllers/usercontroller"

	authFactory "github.com/megalypse/golang-verifymy-backend-test/internal/factory/usecase/auth"
)

var userController controllers.BaseController
var authController controllers.BaseController

func init() {
	userController = usercontroller.UserController{
		CreateUserUsecase:   userUsecaseFactory.GetCreateUserUsecase(),
		UpdateUserUsecase:   userUsecaseFactory.GetUpdateUserUsecase(),
		DeleteUserUsecase:   userUsecaseFactory.GetDeleteUserUsecase(),
		FindUserByIdUsecase: userUsecaseFactory.GetFindUserByIdUsecase(),
	}

	authController = auth.AuthController{
		AuthUserUsecase: authFactory.GetUserEmailAuth(),
	}
}

func GetControllers() []controllers.BaseController {
	return []controllers.BaseController{
		userController,
		authController,
	}
}
