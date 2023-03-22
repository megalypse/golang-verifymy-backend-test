package userrepository

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/megalypse/golang-verifymy-backend-test/internal/data/repository"
	"github.com/megalypse/golang-verifymy-backend-test/internal/domain/models"
	mysqlrepository "github.com/megalypse/golang-verifymy-backend-test/internal/infra/repository/mysql"
	"github.com/megalypse/golang-verifymy-backend-test/internal/infra/repository/mysql/config"
	internal "github.com/megalypse/golang-verifymy-backend-test/internal/infra/repository/mysql/internal"
)

type MySqlUserRepository struct{}

func (MySqlUserRepository) NewConnection(ctx context.Context) repository.Closable {
	newConnection := config.GetMySqlConnection(ctx)

	return mysqlrepository.NewMySqlClosable(ctx, newConnection)
}

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

// TODO: finish implementation (It's late, I'm tired and want to sleep)
func (MySqlUserRepository) FindById(tx repository.Transaction, id int64) (*models.User, *models.CustomError) {
	result, err := tx.Query(`
	SELECT * FROM users
	WHERE id = ?
	`, id)
	if err != nil {
		return nil, err
	}

	_, err = internal.GetMapFromRows(result.(*sql.Rows))
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (MySqlUserRepository) Update(tx repository.Transaction, source *models.User) (*models.User, *models.CustomError) {
	return nil, nil
}
