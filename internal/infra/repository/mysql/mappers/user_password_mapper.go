package mappers

import (
	"database/sql"

	"github.com/megalypse/golang-verifymy-backend-test/internal/domain/customerrors"
	"github.com/megalypse/golang-verifymy-backend-test/internal/domain/models"
)

type SqlUserPasswordMapper models.UserPassword

func GetUserPasswordFromRow(source *sql.Rows) (*models.UserPassword, *models.CustomError) {
	defer source.Close()

	isValid := source.Next()
	if !isValid {
		return nil, customerrors.MakeNotFoundError("No user password retrieved")
	}

	return extractPasswordFromRow(source), nil
}

func extractPasswordFromRow(source *sql.Rows) *models.UserPassword {
	password := models.UserPassword{}

	var createdAt sql.NullTime

	source.Scan(
		&password.Id,
		&password.Password,
		&password.UserId,
		&createdAt,
	)

	if createdAt.Valid {
		password.CreatedAt = &createdAt.Time
	}

	return &password
}
