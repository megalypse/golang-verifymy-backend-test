package addressrepository

import (
	"database/sql"

	"github.com/megalypse/golang-verifymy-backend-test/internal/data/repository"
	"github.com/megalypse/golang-verifymy-backend-test/internal/domain/models"
	mysqlrepository "github.com/megalypse/golang-verifymy-backend-test/internal/infra/repository/mysql"
	internal "github.com/megalypse/golang-verifymy-backend-test/internal/infra/repository/mysql/internal"
	"github.com/megalypse/golang-verifymy-backend-test/internal/infra/repository/mysql/mappers"
)

type MySqlAddressRepository struct {
	mysqlrepository.BaseMySqlRepository
}

func (MySqlAddressRepository) GetAllByUserId(tx repository.Transaction, userId int64) ([]models.Address, *models.CustomError) {
	result, err := tx.Query(`SELECT * FROM addresses WHERE user_id = ?`, userId)
	if err != nil {
		return nil, err
	}

	rawAddressList, err := internal.GetMapFromRows(result.(*sql.Rows))
	if err != nil {
		return nil, err
	}

	addressList := make([]models.Address, 0, len(rawAddressList))
	for _, rawAddress := range rawAddressList {
		address := (&mappers.SqlAddressMapper{}).FromMap(rawAddress).ToAddress()
		addressList = append(addressList, address)
	}

	return addressList, nil
}
