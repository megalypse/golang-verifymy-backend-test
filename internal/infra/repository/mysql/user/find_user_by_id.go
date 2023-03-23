package userrepository

import (
	"database/sql"

	"github.com/megalypse/golang-verifymy-backend-test/internal/data/repository"
	"github.com/megalypse/golang-verifymy-backend-test/internal/domain/models"
	"github.com/megalypse/golang-verifymy-backend-test/internal/infra/repository/mysql/mappers"
)

func (MySqlUserRepository) FindById(tx repository.Transaction, id int64) (*models.User, *models.CustomError) {
	result, err := tx.Query(`
	SELECT * FROM users
	WHERE id = ? AND deleted_at IS NULL
	`, id)
	if err != nil {
		return nil, err
	}

	rows := result.(*sql.Rows)
	user, err := mappers.UserFromRow(rows)
	if err != nil {
		return nil, err
	}

	return user, nil
}
