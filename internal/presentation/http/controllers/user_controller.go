package controllers

import (
	"net/http"

	"github.com/megalypse/golang-verifymy-backend-test/internal/domain/models"
	"github.com/megalypse/golang-verifymy-backend-test/internal/domain/usecases/usecases"
)

type UserController struct {
	createUserUsecase usecases.CreateUser
}

func (uc UserController) createUser(w http.ResponseWriter, r *http.Request) {
	customRequest, err := parseRequest[models.User](r, nil)
	if err != nil {
		writeError(w, err)
		return
	}

	createdUser, err := uc.createUserUsecase.Create(customRequest.Body)
	if err != nil {
		writeError(w, err)
		return
	}

	writeJsonResponse(w, httpResponse{
		HttpStatus: http.StatusOK,
		Message:    "User created successfully",
		Content:    createdUser,
	})
}
