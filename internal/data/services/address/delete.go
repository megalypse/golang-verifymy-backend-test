package address

import (
	"context"

	"github.com/megalypse/golang-verifymy-backend-test/internal/domain/models"
	factory "github.com/megalypse/golang-verifymy-backend-test/internal/factory/repository/mysql"
)

func (as AddressService) Update(ctx context.Context, source *models.Address) (*models.Address, *models.CustomError) {
	conn := factory.NewSqlConnection(ctx)
	defer conn.CloseConnection()

	tx, err := conn.BeginTransaction()
	if err != nil {
		return nil, err
	}

	if err = as.AddressRepository.Update(tx, source); err != nil {
		tx.Rollback()
		return nil, err
	}

	updatedAddress, err := as.AddressRepository.FindById(tx, source.Id)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	if err = tx.Commit(); err != nil {
		tx.Rollback()
		return nil, err
	}

	return updatedAddress, nil
}
