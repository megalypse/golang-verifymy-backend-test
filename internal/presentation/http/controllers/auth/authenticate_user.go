package auth

import (
	"net/http"

	"github.com/megalypse/golang-verifymy-backend-test/internal/domain/models"
	"github.com/megalypse/golang-verifymy-backend-test/internal/presentation/http/controllers"
	localModels "github.com/megalypse/golang-verifymy-backend-test/internal/presentation/http/controllers/auth/models"
)

func (ac AuthController) authenticateUser(w http.ResponseWriter, r *http.Request) {
	request, err := controllers.ParseRequest[localModels.AuthDto](r, nil)
	if err != nil {
		controllers.WriteError(w, err)
		return
	}

	user := &models.User{
		Email: request.Body.Email,
		UserPassword: &models.UserPassword{
			Password: request.Body.Password,
		},
	}

	_, err = ac.AuthUserUsecase.Auth(r.Context(), user)
	if err != nil {
		controllers.WriteError(w, err)
		return
	}

	// TODO: add user roles down below
	token, err := ac.AuthenticationService.GenerateJwtToken(request.Body.Email, []string{})
	if err != nil {
		controllers.WriteError(w, err)
		return
	}

	controllers.WriteJsonResponse(w, controllers.HttpResponse{
		HttpStatus: 200,
		Message:    "User successfully authenticated",
		Content:    token,
	})
}
