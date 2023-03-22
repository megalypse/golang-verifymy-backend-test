package repository

import "github.com/megalypse/golang-verifymy-backend-test/internal/domain/models"

type UserPasswordRepository interface {
	Create(Transaction, *models.UserPassword) (int64, *models.CustomError)
	FindLatestByUserId(Transaction, int64) (*models.UserPassword, *models.CustomError)
}
