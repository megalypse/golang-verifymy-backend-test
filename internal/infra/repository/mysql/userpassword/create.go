package userpasswordrepository

import (
	"database/sql"

	"github.com/megalypse/golang-verifymy-backend-test/internal/data/repository"
	"github.com/megalypse/golang-verifymy-backend-test/internal/domain/models"
	internal "github.com/megalypse/golang-verifymy-backend-test/internal/infra/repository/mysql/internal"
)

func (MySqlUserPasswordRepository) Create(tx repository.Transaction, source *models.UserPassword) (int64, *models.CustomError) {
	result, cErr := tx.Exec(`
	INSERT INTO users_passwords(password_hash, user_id)
	VALUES (?, ?)
	`, source.Password, source.UserId)
	if cErr != nil {
		return 0, cErr
	}

	return internal.GetLastInsertedId(result.(sql.Result))
}
