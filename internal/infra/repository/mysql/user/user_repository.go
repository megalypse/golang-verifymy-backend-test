package userrepository

import (
	"context"
	"database/sql"

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
