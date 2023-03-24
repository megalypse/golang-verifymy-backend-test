package address

import (
	"context"

	"github.com/megalypse/golang-verifymy-backend-test/internal/domain/models"
	factory "github.com/megalypse/golang-verifymy-backend-test/internal/factory/repository/mysql"
)

func (as AddressService) Create(ctx context.Context, source *models.Address) (*models.Address, *models.CustomError) {
	conn := factory.NewSqlConnection(ctx)
	defer conn.CloseConnection()

	tx, err := conn.BeginTransaction()
	if err != nil {
		return nil, err
	}

	addressId, err := as.AddressRepository.Create(tx, source)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	if err = tx.Commit(); err != nil {
		tx.Rollback()
		return nil, err
	}

	tx, err = conn.BeginTransaction()
	address, err := as.AddressRepository.FindById(tx, addressId)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	if err = tx.Commit(); err != nil {
		tx.Rollback()
		return nil, err
	}

	return address, nil
}
