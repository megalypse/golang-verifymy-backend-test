package usercontroller

import (
	"net/http"

	"github.com/megalypse/golang-verifymy-backend-test/internal/presentation/http/controllers"
	"github.com/megalypse/golang-verifymy-backend-test/internal/presentation/http/controllers/internal"
)

func (uc UserController) findUserById(w http.ResponseWriter, r *http.Request) {
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

	user, err := uc.FindUserByIdUsecase.FindById(r.Context(), userId)
	if err != nil {
		controllers.WriteError(w, err)
		return
	}

	controllers.WriteJsonResponse(w, controllers.HttpResponse{
		HttpStatus: http.StatusFound,
		Message:    "User successfully fetched",
		Content:    user,
	})
}
