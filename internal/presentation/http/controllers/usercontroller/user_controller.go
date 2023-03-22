package usercontroller

import (
	"net/http"

	usecases "github.com/megalypse/golang-verifymy-backend-test/internal/domain/usecases/user"
	"github.com/megalypse/golang-verifymy-backend-test/internal/presentation/http/controllers"
)

type UserController struct {
	CreateUserUsecase   usecases.CreateUser
	UpdateUserUsecase   usecases.UpdateUser
	DeleteUserUsecase   usecases.DeleteUser
	FindUserByIdUsecase usecases.FindUserById
}

func (uc UserController) GetHandlers() []controllers.RouteDefinition {
	return []controllers.RouteDefinition{
		{
			Method:       http.MethodPost,
			Route:        "/user",
			HandlingFunc: uc.createUser,
		},
		{
			Method:       http.MethodPut,
			Route:        "/user",
			HandlingFunc: uc.updateUser,
		},
		{
			Method:       http.MethodDelete,
			Route:        "/user/{userId}",
			HandlingFunc: uc.deleteUser,
		},
		{
			Method:       http.MethodGet,
			Route:        "/user/{userId}",
			HandlingFunc: uc.findUserById,
		},
	}
}
