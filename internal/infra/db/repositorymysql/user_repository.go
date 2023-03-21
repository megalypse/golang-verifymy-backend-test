package repositorymysql

import (
	"database/sql"
	"log"

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
	result, err := tx.Exec(`
	INSERT INTO users(name, email, age)
	VALUES (?, ?, ?)
	`, source.Name, source.Email, source.Age)

	sqlResult := result.(sql.Result)
	log.Println(sqlResult.LastInsertId())
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
