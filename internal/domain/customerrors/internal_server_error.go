package customerrors

import (
	"net/http"

	"github.com/megalypse/golang-verifymy-backend-test/internal/domain/models"
)

func MakeInternalServerError(msg string, source error) *models.CustomError {
	return &models.CustomError{
		Code:    http.StatusInternalServerError,
		Message: msg,
		Source:  source,
	}
}
