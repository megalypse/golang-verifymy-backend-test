package usercontroller

import (
	"net/http"

	"github.com/megalypse/golang-verifymy-backend-test/internal/presentation/http/controllers"
	"github.com/megalypse/golang-verifymy-backend-test/internal/presentation/http/controllers/internal"
)

func (uc UserController) deleteUser(w http.ResponseWriter, r *http.Request) {
	userIdParam := "userId"
	request, err := controllers.ParseRequest[internal.Void](r, &[]string{userIdParam})
	if err != nil {
		controllers.WriteError(w, err)
		return
	}

	userId, err := controllers.ParseId(request.Params[userIdParam])
	if err != nil {
		controllers.WriteError(w, err)
		return
	}

	if err = uc.DeleteUserUsecase.Delete(r.Context(), userId); err != nil {
		controllers.WriteError(w, err)
		return
	}

	controllers.WriteJsonResponse(w, controllers.HttpResponse{
		HttpStatus: http.StatusNoContent,
		Message:    "User successfully deleted",
	})
}
