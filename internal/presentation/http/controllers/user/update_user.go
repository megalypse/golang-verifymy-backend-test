package usercontroller

import (
	"net/http"

	"github.com/megalypse/golang-verifymy-backend-test/internal/domain/models"
	"github.com/megalypse/golang-verifymy-backend-test/internal/presentation/http/controllers"
)

func (uc UserController) updateUser(w http.ResponseWriter, r *http.Request) {
	request, err := controllers.ParseRequest[models.User](r, nil)
	if err != nil {
		controllers.WriteError(w, err)
		return
	}

	updatedUser, err := uc.UpdateUserUsecase.Update(r.Context(), request.Body)
	if err != nil {
		controllers.WriteError(w, err)
		return
	}

	controllers.WriteJsonResponse(w, controllers.HttpResponse{
		HttpStatus: http.StatusOK,
		Message:    "User successfully updated",
		Content:    updatedUser,
	})
}
