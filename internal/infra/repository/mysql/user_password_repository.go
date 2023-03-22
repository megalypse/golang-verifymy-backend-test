package repositorymysql

import (
	"database/sql"

	"github.com/megalypse/golang-verifymy-backend-test/internal/data/repository"
	"github.com/megalypse/golang-verifymy-backend-test/internal/domain/customerrors"
	"github.com/megalypse/golang-verifymy-backend-test/internal/domain/models"
	internal "github.com/megalypse/golang-verifymy-backend-test/internal/infra/repository/mysql/internal"
	"github.com/megalypse/golang-verifymy-backend-test/internal/infra/repository/mysql/mappers"
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

func (MySqlUserPasswordRepository) FindLatestByUserId(tx repository.Transaction, userId int64) (*models.UserPassword, *models.CustomError) {
	result, err := tx.Query(`
	SELECT * FROM users_passwords
	WHERE user_id = ?
	ORDER BY created_at DESC
	LIMIT 1
	`, userId)
	if err != nil {
		return nil, err
	}

	rows, err := internal.GetMapFromRows(result.(*sql.Rows))
	if err != nil {
		return nil, err
	}

	if len(rows) < 1 {
		return nil, customerrors.MakeNotFoundError("User not found")
	}

	password := (&mappers.SqlUserPasswordMapper{}).FromMap(rows[0]).ToUserPassword()
	return &password, nil
}
