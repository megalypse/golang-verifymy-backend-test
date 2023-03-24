package auth

import (
	"net/http"

	"github.com/megalypse/golang-verifymy-backend-test/internal/data/services/security"
	"github.com/megalypse/golang-verifymy-backend-test/internal/domain/usecases/auth"
	httputils "github.com/megalypse/golang-verifymy-backend-test/internal/presentation/http"
)

type AuthController struct {
	AuthUserUsecase       auth.UserSignIn
	AuthenticationService security.AuthenticationService
	AuthorizationService  security.AuthorizationService
}

func (ac AuthController) GetHandlers() []httputils.RouteDefinition {
	return []httputils.RouteDefinition{
		{
			Method:       http.MethodPost,
			Route:        "/auth",
			HandlingFunc: ac.authenticateUser,
			Unprotected:  true,
		},
		{
			Method:       http.MethodPost,
			Route:        "/auth/authorize",
			HandlingFunc: ac.authorizeUser,
		},
	}
}
