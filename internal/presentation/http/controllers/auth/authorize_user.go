package auth

import (
	"net/http"

	_ "github.com/megalypse/golang-verifymy-backend-test/internal/domain/models"

	httputils "github.com/megalypse/golang-verifymy-backend-test/internal/presentation/http"
	"github.com/megalypse/golang-verifymy-backend-test/internal/presentation/http/controllers/auth/dto"
)

// @Summary Authorize an user with the received role
// @Tags Auth
// @Success 204 {object} httputils.HttpResponse[string]
// @Failure 500 {object} models.CustomError "Internal Server Error"
// @Param request body dto.AuthorizeUserDto true "Authorizes an user with the received role"
// @Router /auth/authorize [post]
func (ac AuthController) authorizeUser(w http.ResponseWriter, r *http.Request) {
	request, err := httputils.ParseRequest[dto.AuthorizeUserDto](r, nil)
	if err != nil {
		httputils.WriteError(w, err)
		return
	}

	if err = ac.AuthorizationService.AssignRole(r.Context(), request.Body.UserId, request.Body.UserId); err != nil {
		httputils.WriteError(w, err)
		return
	}

	httputils.WriteJsonResponse(w, httputils.HttpResponse[string]{
		HttpStatus: http.StatusNoContent,
		Message:    "Success",
	})
}
