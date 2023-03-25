package httputils

import (
	"encoding/json"
	"net/http"
	"reflect"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/megalypse/golang-verifymy-backend-test/internal/domain/models"
)

func WriteJsonResponse[T any](w http.ResponseWriter, response HttpResponse[T]) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(response.HttpStatus)
	json.NewEncoder(w).Encode(response)
}

func ParseRequest[T any](r *http.Request, params *[]string) (*ParsedRequest[T], *models.CustomError) {
	holder := new(T)

	if reflect.TypeOf(*holder) != reflect.TypeOf(Void{}) {
		if err := json.NewDecoder(r.Body).Decode(holder); err != nil {
			return nil, &models.CustomError{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
				Source:  err,
			}
		}
	}

	paramMap := make(map[string]string)
	if params != nil {
		for _, v := range *params {
			paramMap[v] = GetUrlParam(r, v)
		}
	}

	httpRequest := ParsedRequest[T]{
		Body:   holder,
		Params: paramMap,
	}

	return &httpRequest, nil
}

func GetUrlParam(r *http.Request, key string) string {
	return chi.URLParam(r, key)
}

func WriteError(w http.ResponseWriter, customError *models.CustomError) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(customError.Code)
	json.NewEncoder(w).Encode(HttpResponse[string]{
		HttpStatus: customError.Code,
		Message:    customError.Message,
	})
}

func ParseId(source string) (int64, *models.CustomError) {
	id, err := strconv.ParseInt(source, 10, 64)
	if err != nil {
		return 0, &models.CustomError{
			Code:    500,
			Message: "Failed on parsing id param: " + err.Error(),
			Source:  err,
		}
	}

	return id, nil
}
