package addressrepository

import (
	"database/sql"
	"net/http"

	"github.com/megalypse/golang-verifymy-backend-test/internal/data/repository"
	"github.com/megalypse/golang-verifymy-backend-test/internal/domain/customerrors"
	"github.com/megalypse/golang-verifymy-backend-test/internal/domain/models"
	internal "github.com/megalypse/golang-verifymy-backend-test/internal/infra/repository/mysql/internal"
)

func (MySqlAddressRepository) Delete(tx repository.Transaction, id int64) *models.CustomError {
	result, err := tx.Exec(`
	UPDATE addresses a
		INNER JOIN users u ON a.id = u.id
	SET a.deleted_at = CURRENT_TIMESTAMP
	WHERE a.id = ?
		AND a.deleted_at IS NULL
		AND u.deleted_at IS NULL
	`, id)
	if err != nil {
		return err
	}

	rows, err := internal.GetAffectedRows(result.(sql.Result))
	if err != nil {
		return err
	}

	if rows < 1 {
		return customerrors.MakeNotFoundError("Address not found")
	}

	if rows > 1 {
		return &models.CustomError{
			Code:    http.StatusInternalServerError,
			Message: "More than one got affected in a single target operation",
		}
	}

	return nil
}
