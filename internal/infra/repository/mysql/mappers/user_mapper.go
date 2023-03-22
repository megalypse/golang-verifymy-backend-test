package mappers

import (
	"time"

	"github.com/megalypse/golang-verifymy-backend-test/internal/domain/models"
)

type SqlPersonMapper models.User

func (pm SqlPersonMapper) ToUser() models.User {
	return models.User(pm)
}

func (pm *SqlPersonMapper) FromMap(source map[string]any) *SqlPersonMapper {
	id := source["id"].(int64)
	name := source["name"].(string)
	age := source["age"].(int)
	email := source["email"].(string)

	pm.Id = id
	pm.Name = name
	pm.Age = age
	pm.Email = email

	createdAt, ok := source["created_at"].(time.Time)
	if ok {
		pm.CreatedDate.CreatedAt = &createdAt
	}

	updatedAt, ok := source["updated_at"].(time.Time)
	if ok {
		pm.UpdatedDate.UpdatedAt = &updatedAt
	}

	deletedAt, ok := source["deleted_at"].(time.Time)
	if ok {
		pm.DeletedDate.DeletedAt = &deletedAt
	}

	return pm
}
