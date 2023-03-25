package userrepository

import (
	"database/sql"

	"github.com/go-sql-driver/mysql"
	"github.com/megalypse/golang-verifymy-backend-test/internal/data/repository"
	"github.com/megalypse/golang-verifymy-backend-test/internal/domain/customerrors"
	"github.com/megalypse/golang-verifymy-backend-test/internal/domain/models"
	internal "github.com/megalypse/golang-verifymy-backend-test/internal/infra/repository/mysql/internal"
)

func (MySqlUserRepository) Create(tx repository.Transaction, source *models.User) (int64, *models.CustomError) {
	result, cErr := tx.Exec(`
	INSERT INTO users(name, email, age)
	VALUES (?, ?, ?)
	`, source.Name, source.Email, source.Age)
	if cErr != nil {
		if cErr.Source != nil {
			sqlErr := cErr.Source.(*mysql.MySQLError)
			switch sqlErr.Number {
			case 1062:
				return 0, customerrors.MakeConflictError("Email is already in use", nil)
			default:
				return 0, cErr
			}
		}
		return 0, cErr
	}

	return internal.GetLastInsertedId(result.(sql.Result))
}
