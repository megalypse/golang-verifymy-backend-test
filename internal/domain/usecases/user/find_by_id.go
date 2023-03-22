package usecases

import (
	"context"

	"github.com/megalypse/golang-verifymy-backend-test/internal/domain/models"
)

type FindUserById interface {
	FindById(context.Context, int64) (*models.User, *models.CustomError)
}
