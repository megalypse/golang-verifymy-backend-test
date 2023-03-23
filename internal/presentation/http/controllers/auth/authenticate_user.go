package auth

import (
	"net/http"

	"github.com/megalypse/golang-verifymy-backend-test/internal/domain/models"
	"github.com/megalypse/golang-verifymy-backend-test/internal/presentation/http/controllers"
	localModels "github.com/megalypse/golang-verifymy-backend-test/internal/presentation/http/controllers/auth/models"
	"github.com/megalypse/golang-verifymy-backend-test/internal/presentation/http/controllers/internal"
)

func (ac AuthController) authenticateUser(w http.ResponseWriter, r *http.Request) {
	request, err := internal.ParseRequest[localModels.AuthDto](r, nil)
	if err != nil {
		internal.WriteError(w, err)
		return
	}

	user := &models.User{
		Email: request.Body.Email,
		UserPassword: &models.UserPassword{
			Password: request.Body.Password,
		},
	}

	valid, err := ac.AuthUserUsecase.Auth(r.Context(), user)
	if err != nil {
		internal.WriteError(w, err)
		return
	}

	internal.WriteJsonResponse(w, controllers.HttpResponse{
		HttpStatus: 200,
		Content:    valid,
	})
}
