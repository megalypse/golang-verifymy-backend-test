package repositorymysql

import (
	"github.com/megalypse/golang-verifymy-backend-test/internal/data/repository"
	"github.com/megalypse/golang-verifymy-backend-test/internal/domain/models"
)

type MySqlUserPasswordRepository struct{}

func (MySqlUserPasswordRepository) Create(repository.Transaction, *models.UserPassword) (int64, *models.CustomError) {
	return 0, nil
}

func (MySqlUserPasswordRepository) FindLatestByUserId(repository.Transaction, int64) (*models.UserPassword, *models.CustomError) {
	return nil, nil
}

func (MySqlUserPasswordRepository) FindById(repository.Transaction, int64) (*models.UserPassword, *models.CustomError) {
	return nil, nil
}
