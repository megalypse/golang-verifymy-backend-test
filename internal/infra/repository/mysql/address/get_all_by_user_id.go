package addressrepository

import (
	"database/sql"
	"log"

	"github.com/megalypse/golang-verifymy-backend-test/internal/data/repository"
	"github.com/megalypse/golang-verifymy-backend-test/internal/domain/models"
	"github.com/megalypse/golang-verifymy-backend-test/internal/infra/repository/mysql/mappers"
)

func (MySqlAddressRepository) GetAllByUserId(tx repository.Transaction, userId int64) ([]models.Address, *models.CustomError) {
	log.Println("A")
	result, err := tx.Query(`SELECT * FROM addresses WHERE user_id = ?`, userId)
	if err != nil {
		return nil, err
	}

	rows := result.(*sql.Rows)
	addressList, err := mappers.MapMany(mappers.AddressMapperFunc, rows)
	if err != nil {
		return nil, err
	}

	return addressList, nil
}
