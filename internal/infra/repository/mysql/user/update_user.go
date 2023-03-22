package userrepository

import (
	"database/sql"

	"github.com/megalypse/golang-verifymy-backend-test/internal/data/repository"
	"github.com/megalypse/golang-verifymy-backend-test/internal/domain/customerrors"
	"github.com/megalypse/golang-verifymy-backend-test/internal/domain/models"
	internal "github.com/megalypse/golang-verifymy-backend-test/internal/infra/repository/mysql/internal"
)

func (MySqlUserRepository) Update(tx repository.Transaction, source *models.User) *models.CustomError {
	result, err := tx.Exec(`
	UPDATE users
	SET name = ?, email = ?, age = ?
	WHERE id = ?
	`, source.Name, source.Email, source.Age, source.Id)
	if err != nil {
		return err
	}

	affectedRows, err := internal.GetAffectedRows(result.(sql.Result))
	if err != nil {
		return err
	}

	if affectedRows < 1 {
		return customerrors.MakeNotFoundError("User not found")
	}

	return nil
}
