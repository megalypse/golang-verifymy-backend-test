package usercontroller

import (
	"net/http"

	"github.com/megalypse/golang-verifymy-backend-test/internal/domain/models"
	httputils "github.com/megalypse/golang-verifymy-backend-test/internal/presentation/http"
)

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
