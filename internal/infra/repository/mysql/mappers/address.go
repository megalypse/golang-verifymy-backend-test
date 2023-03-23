package mappers

import (
	"database/sql"

	"github.com/megalypse/golang-verifymy-backend-test/internal/domain/customerrors"
	"github.com/megalypse/golang-verifymy-backend-test/internal/domain/models"
)

func AddressMapperFunc(source *sql.Rows) (*models.Address, *models.CustomError) {
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
