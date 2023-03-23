package repository

import (
	"github.com/megalypse/golang-verifymy-backend-test/internal/domain/models"
)

type UserRepository interface {
	Create(Transaction, *models.User) (int64, *models.CustomError)
	Delete(Transaction, int64) *models.CustomError
	FindById(Transaction, int64) (*models.User, *models.CustomError)
	Update(Transaction, *models.User) *models.CustomError
	FindByEmail(Transaction, string) (*models.User, *models.CustomError)

	baseRepository
}
