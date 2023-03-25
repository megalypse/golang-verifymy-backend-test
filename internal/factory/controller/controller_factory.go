package factory

import (
	serviceFactory "github.com/megalypse/golang-verifymy-backend-test/internal/factory/service"
	addressUsecaseFactory "github.com/megalypse/golang-verifymy-backend-test/internal/factory/usecase/address"
	userUsecaseFactory "github.com/megalypse/golang-verifymy-backend-test/internal/factory/usecase/user"
	httputils "github.com/megalypse/golang-verifymy-backend-test/internal/presentation/http"

	"github.com/megalypse/golang-verifymy-backend-test/internal/presentation/http/controllers/address"
	"github.com/megalypse/golang-verifymy-backend-test/internal/presentation/http/controllers/auth"
	usercontroller "github.com/megalypse/golang-verifymy-backend-test/internal/presentation/http/controllers/user"

	authFactory "github.com/megalypse/golang-verifymy-backend-test/internal/factory/usecase/auth"
)

var userController httputils.BaseController
var authController httputils.BaseController
var addressController httputils.BaseController

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

	addressController = address.AddressController{
		CreateAddressUsecase: addressUsecaseFactory.GetCreateAddressUsecase(),
		UpdateAddressUsecase: addressUsecaseFactory.GetUpdateAddressUsecase(),
		DeleteAddressUsecase: addressUsecaseFactory.GetDeleteAddressUsecase(),
	}

}

func GetControllers() []httputils.BaseController {
	return []httputils.BaseController{
		userController,
		authController,
		addressController,
	}
}
