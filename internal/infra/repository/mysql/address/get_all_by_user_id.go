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
	addresses := make([]models.Address, 0)

	for rows.Next() {
		address, err := mappers.AddressFromRow(rows)
		if err != nil {
			return nil, err
		}

		addresses = append(addresses, *address)
	}

	return addresses, nil
}
