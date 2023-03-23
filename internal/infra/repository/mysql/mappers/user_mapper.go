package mappers

import (
	"database/sql"

	"github.com/megalypse/golang-verifymy-backend-test/internal/domain/customerrors"
	"github.com/megalypse/golang-verifymy-backend-test/internal/domain/models"
)

func UserFromRow(rows *sql.Rows) (*models.User, *models.CustomError) {
	defer rows.Close()

	isValid := rows.Next()
	if !isValid {
		return nil, customerrors.MakeNotFoundError("No users to be returned")
	}

	return extractUserFromRow(rows)
}

func extractUserFromRow(source *sql.Rows) (*models.User, *models.CustomError) {
	user := models.User{}

	var createdAt sql.NullTime
	var updatedAt sql.NullTime
	var deletedAt sql.NullTime

	err := source.Scan(
		&user.Id,
		&user.Name,
		&user.Email,
		&user.Age,
		&createdAt,
		&updatedAt,
		&deletedAt,
	)

	if err != nil {
		return nil, customerrors.MakeInternalServerError(err.Error(), err)
	}

	if createdAt.Valid {
		user.CreatedAt = &createdAt.Time
	}

	if updatedAt.Valid {
		user.UpdatedAt = &updatedAt.Time
	}

	if deletedAt.Valid {
		user.DeletedAt = &deletedAt.Time
	}

	return &user, nil
}
