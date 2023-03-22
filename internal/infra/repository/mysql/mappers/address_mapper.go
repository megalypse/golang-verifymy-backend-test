package mappers

import (
	"time"

	"github.com/megalypse/golang-verifymy-backend-test/internal/domain/models"
)

type SqlAddressMapper models.Address

func (am *SqlAddressMapper) FromMap(source map[string]any) *SqlAddressMapper {
	am.Id = source["id"].(int64)
	am.AddressAlias = source["alias"].(string)
	am.ZipCode = source["zipcode"].(string)
	am.StreetName = source["street_name"].(string)
	am.Number = source["number"].(string)
	am.State = source["state"].(string)
	am.Country = source["country"].(string)
	am.UserId = source["user_id"].(int64)

	createdAt, ok := source["created_at"].(time.Time)
	if ok {
		am.CreatedAt = &createdAt
	}

	updatedAt, ok := source["updated_at"].(time.Time)
	if ok {
		am.UpdatedAt = &updatedAt
	}

	deletedAt, ok := source["deleted_at"].(time.Time)
	if ok {
		am.DeletedAt = &deletedAt
	}

	return am
}

func (am SqlAddressMapper) ToAddress() models.Address {
	return models.Address(am)
}
