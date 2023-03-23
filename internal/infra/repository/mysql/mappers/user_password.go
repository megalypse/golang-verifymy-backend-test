package mappers

import (
	"database/sql"

	"github.com/megalypse/golang-verifymy-backend-test/internal/domain/customerrors"
	"github.com/megalypse/golang-verifymy-backend-test/internal/domain/models"
)

func UserPasswordMapperFunc(source *sql.Rows) (*models.UserPassword, *models.CustomError) {
	password := models.UserPassword{}

	var createdAt sql.NullTime

	err := source.Scan(
		&password.Id,
		&password.Password,
		&password.UserId,
		&createdAt,
	)

	if err != nil {
		return nil, customerrors.MakeInternalServerError(err.Error(), err)
	}

	if createdAt.Valid {
		password.CreatedAt = &createdAt.Time
	}

	return &password, nil
}
