package auth

import (
	"net/http"

	"github.com/megalypse/golang-verifymy-backend-test/internal/domain/models"
	httputils "github.com/megalypse/golang-verifymy-backend-test/internal/presentation/http"
	"github.com/megalypse/golang-verifymy-backend-test/internal/presentation/http/controllers/dto"
)

// @Summary Authenticate user with email and password
// @Tags Auth
// @Success 201 {object} httputils.HttpResponse[string]
// @Failure 500 {object} models.CustomError "Internal Server Error"
// @Param request body dto.AuthDto true "Authenticates user and return a new JWT token"
// @Router /auth [post]
func (ac AuthController) authenticateUser(w http.ResponseWriter, r *http.Request) {
	request, err := httputils.ParseRequest[dto.AuthDto](r, nil)
	if err != nil {
		httputils.WriteError(w, err)
		return
	}

	user := &models.User{
		Email: request.Body.Email,
		UserPassword: &models.UserPassword{
			Password: []byte(request.Body.Password),
		},
	}

	user, err = ac.AuthUserUsecase.SignIn(r.Context(), user)
	if err != nil {
		httputils.WriteError(w, err)
		return
	}

	roles, err := ac.AuthorizationService.GetUserRoles(r.Context(), user.Id)
	if err != nil {
		httputils.WriteError(w, err)
		return
	}

	userRoles := make([]string, 0, len(roles))
	for _, v := range roles {
		userRoles = append(userRoles, v.Alias)
	}

	token, err := ac.AuthenticationService.GenerateJwtToken(request.Body.Email, userRoles)
	if err != nil {
		httputils.WriteError(w, err)
		return
	}

	httputils.WriteJsonResponse(w, httputils.HttpResponse[string]{
		HttpStatus: 200,
		Message:    "User successfully authenticated",
		Content:    token,
	})
}
