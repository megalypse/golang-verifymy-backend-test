package mappers

import (
	"database/sql"

	"github.com/megalypse/golang-verifymy-backend-test/internal/domain/customerrors"
	"github.com/megalypse/golang-verifymy-backend-test/internal/domain/models"
)

type MapperFunc[T any] func(source *sql.Rows) (*T, *models.CustomError)

func MapOne[T any](mapperFunc MapperFunc[T], rows *sql.Rows) (*T, *models.CustomError) {
	defer rows.Close()

	isValid := rows.Next()
	if !isValid {
		return nil, customerrors.MakeNotFoundError("No row to be mapped")
	}

	return mapperFunc(rows)
}

func MapMany[T any](mapperFunc MapperFunc[T], rows *sql.Rows) ([]T, *models.CustomError) {
	defer rows.Close()

	objList := []T{}

	for rows.Next() {
		obj, err := mapperFunc(rows)
		if err != nil {
			return objList, err
		}

		objList = append(objList, *obj)
	}

	return objList, nil
}
