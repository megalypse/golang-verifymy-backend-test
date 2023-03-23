package userpasswordrepository

import (
	"database/sql"

	"github.com/megalypse/golang-verifymy-backend-test/internal/data/repository"
	"github.com/megalypse/golang-verifymy-backend-test/internal/domain/models"
	"github.com/megalypse/golang-verifymy-backend-test/internal/infra/repository/mysql/mappers"
)

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

	rowsResult := result.(*sql.Rows)
	password, err := mappers.GetUserPasswordFromRow(rowsResult)
	if err != nil {
		return nil, err
	}

	return password, nil
}
