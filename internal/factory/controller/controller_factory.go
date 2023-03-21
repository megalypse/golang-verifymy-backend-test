package factory

import (
	userUsecaseFactory "github.com/megalypse/golang-verifymy-backend-test/internal/factory/usecase/user"
	"github.com/megalypse/golang-verifymy-backend-test/internal/presentation/http/controllers"
	"github.com/megalypse/golang-verifymy-backend-test/internal/presentation/http/controllers/usercontroller"
)

var userController controllers.BaseController

func init() {
	userController = usercontroller.UserController{
		CreateUserUsecase: userUsecaseFactory.GetCreateUserUsecase(),
	}
}

func GetControllers() []controllers.BaseController {
	return []controllers.BaseController{
		userController,
	}
}
