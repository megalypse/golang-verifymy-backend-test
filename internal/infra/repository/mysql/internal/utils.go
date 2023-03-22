package internal

import (
	"database/sql"
	"net/http"

	"github.com/megalypse/golang-verifymy-backend-test/internal/domain/models"
)

func GetLastInsertedId(result sql.Result) (int64, *models.CustomError) {
	id, err := result.LastInsertId()
	if err != nil {
		return 0, &models.CustomError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
			Source:  err,
		}
	}

	return id, nil
}

func GetRowsAffected(result sql.Result) (int64, *models.CustomError) {
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, &models.CustomError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
			Source:  err,
		}
	}

	return rowsAffected, nil
}

func GetMapFromRows(rows *sql.Rows) ([]map[string]any, *models.CustomError) {
	cols, _ := rows.Columns()
	rawMaps := make([]map[string]any, 0)

	for rows.Next() {
		columns := make([]any, len(cols))
		columnPointers := make([]any, len(columns))

		for i := range columns {
			columnPointers[i] = &columns[i]
		}

		if err := rows.Scan(columnPointers...); err != nil {
			return rawMaps, &models.CustomError{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
				Source:  err,
			}
		}

		m := make(map[string]any)
		for i, colName := range cols {
			val := columnPointers[i].(*any)
			m[colName] = *val
		}

		rawMaps = append(rawMaps, m)
	}

	return rawMaps, nil
}
