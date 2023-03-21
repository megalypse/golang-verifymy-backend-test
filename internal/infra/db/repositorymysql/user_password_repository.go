package repositorymysql

import (
	"database/sql"
	"net/http"

	"github.com/megalypse/golang-verifymy-backend-test/internal/data/repository"
	"github.com/megalypse/golang-verifymy-backend-test/internal/domain/models"
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

	mySqlResult := result.(sql.Result)
	passwordId, err := mySqlResult.LastInsertId()
	if err != nil {
		return 0, &models.CustomError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
			Source:  err,
		}
	}

	return passwordId, nil
}

func (MySqlUserPasswordRepository) FindLatestByUserId(repository.Transaction, int64) (*models.UserPassword, *models.CustomError) {
	return nil, nil
}

func (MySqlUserPasswordRepository) FindById(repository.Transaction, int64) (*models.UserPassword, *models.CustomError) {
	return nil, nil
}
