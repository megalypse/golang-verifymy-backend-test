package usercontroller

import (
	"net/http"

	_ "github.com/megalypse/golang-verifymy-backend-test/internal/domain/models"
	httputils "github.com/megalypse/golang-verifymy-backend-test/internal/presentation/http"
)

// @Summary Deletes a user
// @Tags Users
// @Success 204
// @Failure 500 {object} models.CustomError "Internal Server Error"
// @Param userId path int true "Person ID"
// @Router /user/{userId} [delete]
// @Security ApiKeyAuth
func (uc UserController) deleteUser(w http.ResponseWriter, r *http.Request) {
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

	if err = uc.DeleteUserUsecase.Delete(r.Context(), userId); err != nil {
		httputils.WriteError(w, err)
		return
	}

	httputils.WriteJsonResponse(w, httputils.HttpResponse[any]{
		HttpStatus: http.StatusOK,
		Message:    "User successfully deleted",
	})
}
