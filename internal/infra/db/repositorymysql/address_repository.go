package repositorymysql

import (
	"github.com/megalypse/golang-verifymy-backend-test/internal/data/repository"
	"github.com/megalypse/golang-verifymy-backend-test/internal/domain/models"
)

type MySqlAddressRepository struct{}

func (MySqlAddressRepository) Create(repository.Transaction, *models.Address) (*models.Address, *models.CustomError) {
	return nil, nil
}

func (MySqlAddressRepository) Delete(repository.Transaction, int64) *models.CustomError {
	return nil
}

func (MySqlAddressRepository) GetLatestByUserId(repository.Transaction, int64) (*models.Address, *models.CustomError) {
	return nil, nil
}

func (MySqlAddressRepository) Update(repository.Transaction, *models.Address) (*models.Address, *models.CustomError) {
	return nil, nil
}
