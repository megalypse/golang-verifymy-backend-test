package mappers

import (
	"time"

	"github.com/megalypse/golang-verifymy-backend-test/internal/domain/models"
)

type SqlUserPasswordMapper models.UserPassword

// TODO: update mapper to use "FromRows" model
func (pm *SqlUserPasswordMapper) FromMap(source map[string]any) *SqlUserPasswordMapper {
	pm.Id = source["id"].(int64)
	pm.Password = source["password_hash"].([]byte)
	pm.UserId = source["user_id"].(int64)

	createdAt, ok := source["created_at"].(time.Time)
	if ok {
		pm.CreatedAt = &createdAt
	}

	return pm
}

func (pm SqlUserPasswordMapper) ToUserPassword() models.UserPassword {
	return models.UserPassword(pm)
}
