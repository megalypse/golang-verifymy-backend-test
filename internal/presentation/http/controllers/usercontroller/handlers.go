package usercontroller

import (
	"net/http"

	"github.com/megalypse/golang-verifymy-backend-test/internal/domain/models"
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
