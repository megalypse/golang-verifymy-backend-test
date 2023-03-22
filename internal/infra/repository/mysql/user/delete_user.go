package userrepository

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/megalypse/golang-verifymy-backend-test/internal/data/repository"
	"github.com/megalypse/golang-verifymy-backend-test/internal/domain/models"
	internal "github.com/megalypse/golang-verifymy-backend-test/internal/infra/repository/mysql/internal"
)

func (MySqlUserRepository) Delete(tx repository.Transaction, id int64) *models.CustomError {
	result, err := tx.Exec(`
	UPDATE users
	SET deleted_at = CURRENT_TIMESTAMP
	WHERE id = ?
	`, id)

	if err != nil {
		return err
	}

	rowsAffected, err := internal.GetRowsAffected(result.(sql.Result))
	if err != nil {
		return err
	}

	if rowsAffected < 1 {
		return &models.CustomError{
			Code:    http.StatusNotFound,
			Message: fmt.Sprintf("No user with id %d found", id),
		}
	}

	return nil
}
