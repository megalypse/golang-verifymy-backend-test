package factory

import (
	serviceFactory "github.com/megalypse/golang-verifymy-backend-test/internal/factory/service"
	userUsecaseFactory "github.com/megalypse/golang-verifymy-backend-test/internal/factory/usecase/user"
	httputils "github.com/megalypse/golang-verifymy-backend-test/internal/presentation/http"
	"github.com/megalypse/golang-verifymy-backend-test/internal/presentation/http/controllers/auth"
	usercontroller "github.com/megalypse/golang-verifymy-backend-test/internal/presentation/http/controllers/user"

	authFactory "github.com/megalypse/golang-verifymy-backend-test/internal/factory/usecase/auth"
)

var userController httputils.BaseController
var authController httputils.BaseController

func init() {
	userController = usercontroller.UserController{
		CreateUserUsecase:   userUsecaseFactory.GetCreateUserUsecase(),
		UpdateUserUsecase:   userUsecaseFactory.GetUpdateUserUsecase(),
		DeleteUserUsecase:   userUsecaseFactory.GetDeleteUserUsecase(),
		FindUserByIdUsecase: userUsecaseFactory.GetFindUserByIdUsecase(),
	}

	authController = auth.AuthController{
		AuthUserUsecase:       authFactory.GetUserEmailAuth(),
		AuthenticationService: serviceFactory.GetJwtAuthenticationService(),
		AuthorizationService:  serviceFactory.GetRolesAuthorizationService(),
	}
}

func GetControllers() []httputils.BaseController {
	return []httputils.BaseController{
		userController,
		authController,
	}
}

func GetUserController() httputils.BaseController {
	return userController
}
