package repositorymysql

import (
	"github.com/megalypse/golang-verifymy-backend-test/internal/data/repository"
	"github.com/megalypse/golang-verifymy-backend-test/internal/domain/models"
	"github.com/megalypse/golang-verifymy-backend-test/internal/infra/db/repositorymysql/config"
)

type MySqlUserRepository struct{}

func (MySqlUserRepository) NewConnection() repository.Connectable {
	return &MySqlTransactionable{
		connection: config.GetMySqlConnection(),
	}
}

func (MySqlUserRepository) Create(tx repository.Transaction, source *models.User) (int64, *models.CustomError) {
	_, err := tx.Exec()
	if err != nil {
		return 0, err
	}

	return 0, nil
}

func (MySqlUserRepository) Delete(tx repository.Transaction, id int64) *models.CustomError {
	return nil
}

func (MySqlUserRepository) FindById(tx repository.Transaction, id int64) (*models.User, *models.CustomError) {
	return nil, nil
}

func (MySqlUserRepository) Update(tx repository.Transaction, source *models.User) (*models.User, *models.CustomError) {
	return nil, nil
}
