package repository

import "github.com/megalypse/golang-verifymy-backend-test/internal/domain/models"

type AddressRepository interface {
	Create(Transaction, *models.Address) (int64, *models.CustomError)
	Delete(Transaction, int64) *models.CustomError
	GetAllByUserId(Transaction, int64) ([]models.Address, *models.CustomError)
	Update(Transaction, *models.Address) *models.CustomError
}
