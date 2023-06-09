package service

import (
	"context"

	"github.com/megalypse/golang-verifymy-backend-test/internal/domain/models"
	factory "github.com/megalypse/golang-verifymy-backend-test/internal/factory/repository/mysql"
)

func (us UserService) Delete(ctx context.Context, id int64) *models.CustomError {
	connection := factory.NewSqlConnection(ctx)
	defer connection.CloseConnection()

	tx, err := connection.BeginTransaction()
	if err != nil {
		return err
	}

	if err := us.userRepository.Delete(tx, id); err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}
