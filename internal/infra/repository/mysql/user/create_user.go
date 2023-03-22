package userrepository

import (
	"database/sql"

	"github.com/megalypse/golang-verifymy-backend-test/internal/data/repository"
	"github.com/megalypse/golang-verifymy-backend-test/internal/domain/models"
	internal "github.com/megalypse/golang-verifymy-backend-test/internal/infra/repository/mysql/internal"
)

func (MySqlUserRepository) Create(tx repository.Transaction, source *models.User) (int64, *models.CustomError) {
	result, cErr := tx.Exec(`
	INSERT INTO users(name, email, age)
	VALUES (?, ?, ?)
	`, source.Name, source.Email, source.Age)
	if cErr != nil {
		return 0, cErr
	}

	return internal.GetLastInsertedId(result.(sql.Result))
}
