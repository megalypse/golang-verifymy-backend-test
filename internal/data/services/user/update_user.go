package service

import (
	"context"

	"github.com/megalypse/golang-verifymy-backend-test/internal/domain/models"
	factory "github.com/megalypse/golang-verifymy-backend-test/internal/factory/repository/mysql"
)

func (us UserService) Update(ctx context.Context, source *models.User) (*models.User, *models.CustomError) {
	connection := factory.NewSqlConnection(ctx)
	defer connection.CloseConnection()

	tx, err := connection.BeginTransaction()
	if err != nil {
		return nil, err
	}

	err = us.userRepository.Update(tx, source)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	if err = tx.Commit(); err != nil {
		tx.Rollback()
		return nil, err
	}

	user, err := us.findById(ctx, connection, source.Id)
	if err != nil {
		return nil, err
	}

	return user, nil
}
