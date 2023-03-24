package usercontroller

import (
	"net/http"

	"github.com/megalypse/golang-verifymy-backend-test/internal/domain/models"
	httputils "github.com/megalypse/golang-verifymy-backend-test/internal/presentation/http"
)

// @Summary Creates a new user
// @Tags Users
// @Success 201 {object} models.User
// @Failure 422 {object} models.CustomError "Unprocessable Entity"
// @Failure 500 {object} models.CustomError "Internal Server Error"
// @Param request body models.User true "Create user request"
// @Router /user [post]
func (uc UserController) createUser(w http.ResponseWriter, r *http.Request) {
	customRequest, err := httputils.ParseRequest[models.User](r, nil)
	if err != nil {
		httputils.WriteError(w, err)
		return
	}

	createdUser, err := uc.CreateUserUsecase.Create(r.Context(), customRequest.Body)
	if err != nil {
		httputils.WriteError(w, err)
		return
	}

	httputils.WriteJsonResponse(w, httputils.HttpResponse{
		HttpStatus: http.StatusOK,
		Message:    "User created successfully",
		Content:    createdUser,
	})
}
