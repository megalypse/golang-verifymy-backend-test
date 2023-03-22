package repositorymysql

import (
	"database/sql"

	"github.com/megalypse/golang-verifymy-backend-test/internal/data/repository"
	"github.com/megalypse/golang-verifymy-backend-test/internal/domain/models"
	internal "github.com/megalypse/golang-verifymy-backend-test/internal/infra/repository/mysql/internal"
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

	return internal.GetLastInsertedId(result.(sql.Result))
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
