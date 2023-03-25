package usecases

import (
	"context"

	"github.com/megalypse/golang-verifymy-backend-test/internal/domain/models"
)

type CreateUser interface {
	Create(context.Context, *models.User) (*models.User, *models.CustomError)
}
