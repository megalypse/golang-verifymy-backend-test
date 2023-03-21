package repositorymysql

import "github.com/megalypse/golang-verifymy-backend-test/internal/domain/models"

type MySqlUserRepository struct{}

func (MySqlUserRepository) Create(*models.User) (*models.User, *models.CustomError) {
	return nil, nil
}

func (MySqlUserRepository) Delete(int64) *models.CustomError {
	return nil
}

func (MySqlUserRepository) FindById(int64) (*models.User, *models.CustomError) {
	return nil, nil
}

func (MySqlUserRepository) Update(*models.User) (*models.User, *models.CustomError) {
	return nil, nil
}
