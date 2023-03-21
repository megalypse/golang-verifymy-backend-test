package repositorymysql

import (
	"database/sql"
	"net/http"

	"github.com/megalypse/golang-verifymy-backend-test/internal/data/repository"
	"github.com/megalypse/golang-verifymy-backend-test/internal/domain/models"
)

type MySqlAddressRepository struct{}

func (MySqlAddressRepository) Create(tx repository.Transaction, source *models.Address) (int64, *models.CustomError) {
	result, cErr := tx.Exec(`
	INSERT INTO addresses(alias, zipcode, street_name, number, state, country, user_id)
	VALUES (?, ?, ?, ?, ?, ?, ?)
	`,
		source.AddressAlias,
		source.ZipCode,
		source.StreetName,
		source.Number,
		source.State,
		source.Country,
		source.UserId,
	)
	if cErr != nil {
		return 0, cErr
	}

	mySqlResult := result.(sql.Result)
	addressId, err := mySqlResult.LastInsertId()
	if err != nil {
		return 0, &models.CustomError{
			Code:    http.StatusInternalServerError,
			Message: "Failed on saving address.\n" + err.Error(),
			Source:  err,
		}
	}

	return addressId, nil
}

func (MySqlAddressRepository) Delete(repository.Transaction, int64) *models.CustomError {
	return nil
}

func (MySqlAddressRepository) GetAllByUserId(repository.Transaction, int64) ([]models.Address, *models.CustomError) {
	return nil, nil
}

func (MySqlAddressRepository) Update(repository.Transaction, *models.Address) (*models.Address, *models.CustomError) {
	return nil, nil
}
