package service

import (
	"context"

	"github.com/megalypse/golang-verifymy-backend-test/internal/data/repository"
	"github.com/megalypse/golang-verifymy-backend-test/internal/domain/models"
	factory "github.com/megalypse/golang-verifymy-backend-test/internal/factory/repository/mysql"
)

func (us UserService) FindById(ctx context.Context, id int64) (*models.User, *models.CustomError) {
	connection := factory.NewSqlConnection(ctx)
	defer connection.CloseConnection()

	return us.findById(ctx, connection, id)
}

func (us UserService) findById(ctx context.Context, conn repository.Closable, id int64) (*models.User, *models.CustomError) {
	tx, err := conn.BeginTransaction()
	if err != nil {
		return nil, err
	}

	result, err := us.userRepository.FindById(tx, id)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()
	return result, nil
}
