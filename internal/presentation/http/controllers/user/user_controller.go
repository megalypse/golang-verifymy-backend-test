package usercontroller

import (
	"net/http"

	usecases "github.com/megalypse/golang-verifymy-backend-test/internal/domain/usecases/user"
	httputils "github.com/megalypse/golang-verifymy-backend-test/internal/presentation/http"
	"github.com/megalypse/golang-verifymy-backend-test/internal/presentation/http/controllers/roles"
)

type UserController struct {
	CreateUserUsecase   usecases.CreateUser
	UpdateUserUsecase   usecases.UpdateUser
	DeleteUserUsecase   usecases.DeleteUser
	FindUserByIdUsecase usecases.FindUserById
}

func (uc UserController) GetHandlers() []httputils.RouteDefinition {
	return []httputils.RouteDefinition{
		{
			Method:       http.MethodPost,
			Route:        "/user",
			HandlingFunc: uc.createUser,
			Unprotected:  true,
		},
		{
			Method:        http.MethodPut,
			Route:         "/user",
			HandlingFunc:  uc.updateUser,
			RequiredRoles: []string{roles.UPDATE, roles.READ},
		},
		{
			Method:        http.MethodDelete,
			Route:         "/user/{userId}",
			HandlingFunc:  uc.deleteUser,
			RequiredRoles: []string{roles.DELETE},
		},
		{
			Method:        http.MethodGet,
			Route:         "/user/{userId}",
			HandlingFunc:  uc.findUserById,
			RequiredRoles: []string{roles.READ},
		},
	}
}
