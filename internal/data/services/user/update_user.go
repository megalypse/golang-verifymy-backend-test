package service

import (
	"context"

	"github.com/megalypse/golang-verifymy-backend-test/internal/domain/models"
)

func (us UserService) Update(ctx context.Context, source *models.User) (*models.User, *models.CustomError) {
	connection := us.userRepository.NewConnection(ctx)
	defer connection.CloseConnection()

	tx, err := connection.BeginTransaction()
	if err != nil {
		return nil, err
	}

	result, err := us.userRepository.Update(tx, source)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()
	return result, nil
}
