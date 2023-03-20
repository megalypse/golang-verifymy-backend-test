package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/megalypse/golang-verifymy-backend-test/internal/domain/models"
)

func writeJsonResponse(w http.ResponseWriter, response httpResponse) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(response.HttpStatus)
	json.NewEncoder(w).Encode(response)
}

func parseRequest[T any](r *http.Request, params *[]string) (*parsedRequest[T], *models.CustomError) {
	holder := new(T)
	if err := json.NewDecoder(r.Body).Decode(holder); err != nil {
		return nil, &models.CustomError{
			Code:    http.StatusInternalServerError,
			Message: "Failed on parsing request body",
			Source:  err,
		}
	}

	paramMap := make(map[string]string)
	if params != nil {
		for _, v := range *params {
			paramMap[v] = getUrlParam(r, v)
		}
	}

	httpRequest := parsedRequest[T]{
		Body:   holder,
		Params: paramMap,
	}

	return &httpRequest, nil
}

func getUrlParam(r *http.Request, key string) string {
	return chi.URLParam(r, key)
}

func writeError(w http.ResponseWriter, customError *models.CustomError) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(customError.Code)
	json.NewEncoder(w).Encode(customError)
}
