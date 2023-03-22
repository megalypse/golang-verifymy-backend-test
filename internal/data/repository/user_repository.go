package repository

import (
	"context"

	"github.com/megalypse/golang-verifymy-backend-test/internal/domain/models"
)

type UserRepository interface {
	Create(Transaction, *models.User) (int64, *models.CustomError)
	Delete(Transaction, int64) *models.CustomError
	FindById(Transaction, int64) (*models.User, *models.CustomError)
	Update(Transaction, *models.User) (*models.User, *models.CustomError)
	NewConnection(context.Context) Closable
}
