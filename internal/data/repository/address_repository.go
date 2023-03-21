package repository

import "github.com/megalypse/golang-verifymy-backend-test/internal/domain/models"

type AddressRepository interface {
	Create(Transaction, *models.Address) (*models.Address, *models.CustomError)
	Delete(Transaction, int64) *models.CustomError
	GetLatestByUserId(Transaction, int64) (*models.Address, *models.CustomError)
	Update(Transaction, *models.Address) (*models.Address, *models.CustomError)
}
