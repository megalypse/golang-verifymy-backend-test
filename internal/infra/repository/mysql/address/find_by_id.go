package addressrepository

import (
	"database/sql"

	"github.com/megalypse/golang-verifymy-backend-test/internal/data/repository"
	"github.com/megalypse/golang-verifymy-backend-test/internal/domain/models"
	"github.com/megalypse/golang-verifymy-backend-test/internal/infra/repository/mysql/mappers"
)

func (MySqlAddressRepository) FindById(tx repository.Transaction, addressId int64) (*models.Address, *models.CustomError) {
	result, err := tx.Query(`SELECT * FROM addresses WHERE id = ? LIMIT 1`, addressId)
	if err != nil {
		return nil, err
	}

	rows := result.(*sql.Rows)
	address, err := mappers.MapOne(mappers.AddressMapperFunc, rows)
	if err != nil {
		return nil, err
	}

	return address, nil
}
