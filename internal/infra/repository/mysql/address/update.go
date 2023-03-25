package addressrepository

import (
	"database/sql"

	"github.com/megalypse/golang-verifymy-backend-test/internal/data/repository"
	"github.com/megalypse/golang-verifymy-backend-test/internal/domain/customerrors"
	"github.com/megalypse/golang-verifymy-backend-test/internal/domain/models"
	internal "github.com/megalypse/golang-verifymy-backend-test/internal/infra/repository/mysql/internal"
)

func (MySqlAddressRepository) Update(tx repository.Transaction, source *models.Address) *models.CustomError {
	result, err := tx.Exec(`
	UPDATE addresses
	SET alias = ?, zipcode = ?, street_name = ?, number = ?, state = ?, country = ?
	WHERE id = ? AND deleted_at IS NULL
	`, source.AddressAlias, source.ZipCode, source.StreetName, source.Number, source.State, source.Country, source.Id)
	if err != nil {
		return err
	}

	rows, err := internal.GetAffectedRows(result.(sql.Result))
	if err != nil {
		return err
	}

	if rows < 1 {
		return customerrors.MakeNotFoundError("Address not found")
	}

	if rows > 1 {
		return customerrors.MakeConflictError("More than one row affected", nil)
	}

	return nil
}
