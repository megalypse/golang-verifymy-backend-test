package mappers

import (
	"database/sql"

	"github.com/megalypse/golang-verifymy-backend-test/internal/domain/customerrors"
	"github.com/megalypse/golang-verifymy-backend-test/internal/domain/models"
)

func AddressFromRow(rows *sql.Rows) (*models.Address, *models.CustomError) {
	defer rows.Close()

	isValid := rows.Next()

	if !isValid {
		return nil, customerrors.MakeNotFoundError("No address to be returned")
	}

	return addressFromRow(rows)
}

func ManyAddressesFromRows(rows *sql.Rows) ([]models.Address, *models.CustomError) {
	defer rows.Close()

	addressList := []models.Address{}

	for rows.Next() {
		address, err := addressFromRow(rows)
		if err != nil {
			return addressList, err
		}

		addressList = append(addressList, *address)
	}

	return addressList, nil
}

func addressFromRow(source *sql.Rows) (*models.Address, *models.CustomError) {
	address := models.Address{}
	var createdAt sql.NullTime
	var updatedAt sql.NullTime
	var deletedAt sql.NullTime

	err := source.Scan(
		&address.Id,
		&address.AddressAlias,
		&address.ZipCode,
		&address.StreetName,
		&address.Number,
		&address.State,
		&address.Country,
		&address.UserId,
		&createdAt,
		&updatedAt,
		&deletedAt,
	)

	if err != nil {
		return nil, customerrors.MakeInternalServerError(err.Error(), err)
	}

	if createdAt.Valid {
		address.CreatedAt = &createdAt.Time
	}

	if updatedAt.Valid {
		address.UpdatedAt = &updatedAt.Time
	}

	if deletedAt.Valid {
		address.DeletedAt = &deletedAt.Time
	}

	return &address, nil
}
