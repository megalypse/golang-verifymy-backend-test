package usecases

import (
	"context"

	"github.com/megalypse/golang-verifymy-backend-test/internal/domain/models"
)

type UpdateUser interface {
	Update(context.Context, *models.User) (*models.User, *models.CustomError)
}
