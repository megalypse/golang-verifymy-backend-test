package repositorymysql

import (
	"database/sql"

	"github.com/megalypse/golang-verifymy-backend-test/internal/data/repository"
	"github.com/megalypse/golang-verifymy-backend-test/internal/domain/models"
	internal "github.com/megalypse/golang-verifymy-backend-test/internal/infra/repository/mysql/internal"
)

type MySqlUserPasswordRepository struct{}

func (MySqlUserPasswordRepository) Create(tx repository.Transaction, source *models.UserPassword) (int64, *models.CustomError) {
	result, cErr := tx.Exec(`
	INSERT INTO users_passwords(password_hash, salt, user_id)
	VALUES (?, ?, ?)
	`, source.Password, source.Salt, source.UserId)
	if cErr != nil {
		return 0, cErr
	}

	return internal.GetLastInsertedId(result.(sql.Result))
}

func (MySqlUserPasswordRepository) FindLatestByUserId(repository.Transaction, int64) (*models.UserPassword, *models.CustomError) {
	return nil, nil
}

func (MySqlUserPasswordRepository) FindById(repository.Transaction, int64) (*models.UserPassword, *models.CustomError) {
	return nil, nil
}
