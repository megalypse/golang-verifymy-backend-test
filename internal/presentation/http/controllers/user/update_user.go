package usercontroller

import (
	"net/http"

	"github.com/megalypse/golang-verifymy-backend-test/internal/domain/models"
	httputils "github.com/megalypse/golang-verifymy-backend-test/internal/presentation/http"
)

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
