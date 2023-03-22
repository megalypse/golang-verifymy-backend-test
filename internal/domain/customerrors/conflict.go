package customerrors

import (
	"net/http"

	"github.com/megalypse/golang-verifymy-backend-test/internal/domain/models"
)

func MakeConflictError(message string) *models.CustomError {
	return &models.CustomError{
		Code:    http.StatusConflict,
		Message: message,
	}
}
