package mappers

import (
	"database/sql"

	"github.com/megalypse/golang-verifymy-backend-test/internal/domain/customerrors"
	"github.com/megalypse/golang-verifymy-backend-test/internal/domain/models"
)

func UserRoleMapperFunc(rows *sql.Rows) (*models.UserRole, *models.CustomError) {
	role := models.UserRole{}

	err := rows.Scan(
		&role.UserId,
		&role.RoleId,
	)

	if err != nil {
		return nil, customerrors.MakeInternalServerError(err.Error(), err)
	}

	return &role, nil
}
