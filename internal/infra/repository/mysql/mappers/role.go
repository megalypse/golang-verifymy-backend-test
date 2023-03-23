package mappers

import (
	"database/sql"

	"github.com/megalypse/golang-verifymy-backend-test/internal/domain/customerrors"
	"github.com/megalypse/golang-verifymy-backend-test/internal/domain/models"
)

func RoleMapperFunc(rows *sql.Rows) (*models.Role, *models.CustomError) {
	role := models.Role{}

	var createdAt sql.NullTime

	err := rows.Scan(
		&role.Id,
		&role.Alias,
		&createdAt,
	)

	if err != nil {
		return nil, customerrors.MakeInternalServerError(err.Error(), err)
	}

	if createdAt.Valid {
		role.CreatedAt = &createdAt.Time
	}

	return &role, nil
}
