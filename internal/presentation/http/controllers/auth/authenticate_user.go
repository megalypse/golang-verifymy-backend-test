package auth

import (
	"log"
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

	user, err = ac.AuthUserUsecase.SignIn(r.Context(), user)
	if err != nil {
		controllers.WriteError(w, err)
		return
	}

	roles, err := ac.AuthorizationService.GetUserRoles(r.Context(), user.Id)
	if err != nil {
		controllers.WriteError(w, err)
		return
	}

	userRoles := make([]string, 0, len(roles))
	for _, v := range roles {
		userRoles = append(userRoles, v.Alias)
	}

	log.Println(userRoles)

	token, err := ac.AuthenticationService.GenerateJwtToken(request.Body.Email, userRoles)
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
