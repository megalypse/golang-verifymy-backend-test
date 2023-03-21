package repositorymysql

import (
	"database/sql"
	"net/http"

	"github.com/megalypse/golang-verifymy-backend-test/internal/data/repository"
	"github.com/megalypse/golang-verifymy-backend-test/internal/domain/models"
)

type MySqlTransactionable struct {
	connection *sql.DB
}

func (mst *MySqlTransactionable) CloseConnection() {
	mst.connection.Close()
}

func (mst MySqlTransactionable) BeginTransaction() (repository.Transaction, *models.CustomError) {
	tx, err := mst.connection.Begin()
	if err != nil {
		return nil, &models.CustomError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
			Source:  err,
		}
	}

	transaction := MySqlTransaction{
		tx: tx,
	}

	return &transaction, nil
}

type MySqlTransaction struct {
	tx *sql.Tx
}

func (mst *MySqlTransaction) Rollback() *models.CustomError {
	if err := mst.tx.Rollback(); err != nil {
		return &models.CustomError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
			Source:  err,
		}
	}

	return nil
}

func (mst *MySqlTransaction) Commit() *models.CustomError {
	if err := mst.tx.Commit(); err != nil {
		return &models.CustomError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
			Source:  err,
		}
	}

	return nil
}

func (mst *MySqlTransaction) Exec(args ...any) (any, *models.CustomError) {
	rawSttmt := args[0]

	switch sttmt := rawSttmt.(type) {
	case string:
		sqlArgs := args[1:]

		result, err := mst.tx.Exec(sttmt, sqlArgs...)
		if err != nil {
			return nil, &models.CustomError{
				Code:    http.StatusInternalServerError,
				Message: "Failed on executing transaction",
				Source:  err,
			}
		}

		id, err := result.LastInsertId()
		if err != nil {
			return nil, &models.CustomError{
				Code:    http.StatusInternalServerError,
				Message: "Failed on getting last inserted ID",
				Source:  err,
			}
		}

		return id, nil
	default:
		return nil, &models.CustomError{
			Code:    http.StatusBadRequest,
			Message: "Invalid type for query. Must be a string",
		}
	}
}
