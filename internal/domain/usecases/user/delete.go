package usecases

import (
	"context"

	"github.com/megalypse/golang-verifymy-backend-test/internal/domain/models"
)

type DeleteUser interface {
	Delete(context.Context, int64) *models.CustomError
}
