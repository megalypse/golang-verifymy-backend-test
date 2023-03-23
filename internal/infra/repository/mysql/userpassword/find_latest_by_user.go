package userpasswordrepository

import (
	"database/sql"

	"github.com/megalypse/golang-verifymy-backend-test/internal/data/repository"
	"github.com/megalypse/golang-verifymy-backend-test/internal/domain/customerrors"
	"github.com/megalypse/golang-verifymy-backend-test/internal/domain/models"
	internal "github.com/megalypse/golang-verifymy-backend-test/internal/infra/repository/mysql/internal"
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
	rows, err := internal.GetMapFromRows(rowsResult)
	defer rowsResult.Close()

	if err != nil {
		return nil, err
	}

	if len(rows) < 1 {
		return nil, customerrors.MakeNotFoundError("User not found")
	}

	password := (&mappers.SqlUserPasswordMapper{}).FromMap(rows[0]).ToUserPassword()
	return &password, nil
}
