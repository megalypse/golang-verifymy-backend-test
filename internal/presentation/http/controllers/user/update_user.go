package usercontroller

import (
	"net/http"

	"github.com/megalypse/golang-verifymy-backend-test/internal/domain/models"
	httputils "github.com/megalypse/golang-verifymy-backend-test/internal/presentation/http"
)

// @Summary Updates an user with the received data
// @Tags Users
// @Success 200 {object} models.User
// @Failure 422 {object} models.CustomError "Unprocessable Entity"
// @Failure 500 {object} models.CustomError "Internal Server Error"
// @Param request body models.User true "User model"
// @Router /user [put]
// @Security ApiKeyAuth
func (uc UserController) updateUser(w http.ResponseWriter, r *http.Request) {
	request, err := httputils.ParseRequest[models.User](r, nil)
	if err != nil {
		httputils.WriteError(w, err)
		return
	}

	updatedUser, err := uc.UpdateUserUsecase.Update(r.Context(), request.Body)
	if err != nil {
		httputils.WriteError(w, err)
		return
	}

	httputils.WriteJsonResponse(w, httputils.HttpResponse{
		HttpStatus: http.StatusOK,
		Message:    "User successfully updated",
		Content:    updatedUser,
	})
}
