package userrepository

import (
	"database/sql"

	"github.com/megalypse/golang-verifymy-backend-test/internal/data/repository"
	"github.com/megalypse/golang-verifymy-backend-test/internal/domain/customerrors"
	"github.com/megalypse/golang-verifymy-backend-test/internal/domain/models"
	"github.com/megalypse/golang-verifymy-backend-test/internal/infra/repository/mysql/mappers"
)

func (MySqlUserRepository) FindByEmail(tx repository.Transaction, email string) (*models.User, *models.CustomError) {
	result, err := tx.Query(`
	SELECT * FROM users
	WHERE email = ? AND deleted_at IS NULL
	LIMIT 1
	`, email)
	if err != nil {
		return nil, err
	}

	rows := result.(*sql.Rows)
	defer rows.Close()

	haveNext := rows.Next()

	if !haveNext {
		return nil, customerrors.MakeNotFoundError("User not found")
	}

	user, err := mappers.UserFromRow(rows)
	if err != nil {
		return nil, err
	}

	return user, nil
}
