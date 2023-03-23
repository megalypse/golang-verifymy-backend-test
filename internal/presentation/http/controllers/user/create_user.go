package usercontroller

import (
	"net/http"

	"github.com/megalypse/golang-verifymy-backend-test/internal/domain/models"
	"github.com/megalypse/golang-verifymy-backend-test/internal/presentation/http/controllers"
)

func (uc UserController) createUser(w http.ResponseWriter, r *http.Request) {
	customRequest, err := controllers.ParseRequest[models.User](r, nil)
	if err != nil {
		controllers.WriteError(w, err)
		return
	}

	createdUser, err := uc.CreateUserUsecase.Create(r.Context(), customRequest.Body)
	if err != nil {
		controllers.WriteError(w, err)
		return
	}

	controllers.WriteJsonResponse(w, controllers.HttpResponse{
		HttpStatus: http.StatusOK,
		Message:    "User created successfully",
		Content:    createdUser,
	})
}
