package auth

import (
	"net/http"

	"github.com/megalypse/golang-verifymy-backend-test/internal/domain/usecases/auth"
	"github.com/megalypse/golang-verifymy-backend-test/internal/presentation/http/controllers"
)

type AuthController struct {
	AuthUserUsecase auth.AuthUser
}

func (ac AuthController) GetHandlers() []controllers.RouteDefinition {
	return []controllers.RouteDefinition{
		{
			Method:       http.MethodPost,
			Route:        "/auth",
			HandlingFunc: ac.authenticateUser,
		},
	}
}
