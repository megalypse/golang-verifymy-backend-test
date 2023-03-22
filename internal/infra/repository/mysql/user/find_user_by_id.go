package userrepository

import (
	"database/sql"

	"github.com/megalypse/golang-verifymy-backend-test/internal/data/repository"
	"github.com/megalypse/golang-verifymy-backend-test/internal/domain/customerrors"
	"github.com/megalypse/golang-verifymy-backend-test/internal/domain/models"
	internal "github.com/megalypse/golang-verifymy-backend-test/internal/infra/repository/mysql/internal"
	"github.com/megalypse/golang-verifymy-backend-test/internal/infra/repository/mysql/mappers"
)

func (MySqlUserRepository) FindById(tx repository.Transaction, id int64) (*models.User, *models.CustomError) {
	result, err := tx.Query(`
	SELECT * FROM users
	WHERE id = ?
	`, id)
	if err != nil {
		return nil, err
	}

	rowsMap, err := internal.GetMapFromRows(result.(*sql.Rows))
	if err != nil {
		return nil, err
	}

	if len(rowsMap) < 1 {
		return nil, customerrors.MakeNotFoundError("User not found")
	}

	user := (&mappers.SqlPersonMapper{}).FromMap(rowsMap[0]).ToUser()
	return &user, nil
}
