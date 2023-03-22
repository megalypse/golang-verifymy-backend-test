package customerrors

import (
	"net/http"

	"github.com/megalypse/golang-verifymy-backend-test/internal/domain/models"
)

func MakeNotFoundError(message string) *models.CustomError {
	return &models.CustomError{
		Code:    http.StatusNotFound,
		Message: message,
	}
}
