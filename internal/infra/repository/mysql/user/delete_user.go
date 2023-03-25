package userrepository

import (
	"database/sql"

	"github.com/megalypse/golang-verifymy-backend-test/internal/data/repository"
	"github.com/megalypse/golang-verifymy-backend-test/internal/domain/customerrors"
	"github.com/megalypse/golang-verifymy-backend-test/internal/domain/models"
	internal "github.com/megalypse/golang-verifymy-backend-test/internal/infra/repository/mysql/internal"
)

func (MySqlUserRepository) Delete(tx repository.Transaction, id int64) *models.CustomError {
	result, err := tx.Exec(`
	UPDATE users
	SET deleted_at = CURRENT_TIMESTAMP
	WHERE id = ? AND deleted_at IS NULL
	`, id)

	if err != nil {
		return err
	}

	rowsAffected, err := internal.GetAffectedRows(result.(sql.Result))
	if err != nil {
		return err
	}

	if rowsAffected < 1 {
		return customerrors.MakeNotFoundError("User not found")
	}

	return nil
}
