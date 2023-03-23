package usercontroller

import (
	"net/http"

	httputils "github.com/megalypse/golang-verifymy-backend-test/internal/presentation/http"
)

func (uc UserController) findUserById(w http.ResponseWriter, r *http.Request) {
	userIdParam := "userId"

	request, err := httputils.ParseRequest[httputils.Void](r, &[]string{userIdParam})
	if err != nil {
		httputils.WriteError(w, err)
		return
	}

	userId, err := httputils.ParseId(request.Params[userIdParam])
	if err != nil {
		httputils.WriteError(w, err)
		return
	}

	user, err := uc.FindUserByIdUsecase.FindById(r.Context(), userId)
	if err != nil {
		httputils.WriteError(w, err)
		return
	}

	httputils.WriteJsonResponse(w, httputils.HttpResponse{
		HttpStatus: http.StatusFound,
		Message:    "User successfully fetched",
		Content:    user,
	})
}
