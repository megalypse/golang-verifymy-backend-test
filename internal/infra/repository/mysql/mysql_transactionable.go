package repositorymysql

import (
	"context"
	"database/sql"
	"log"
	"net/http"

	"github.com/megalypse/golang-verifymy-backend-test/internal/data/repository"
	"github.com/megalypse/golang-verifymy-backend-test/internal/domain/customerrors"
	"github.com/megalypse/golang-verifymy-backend-test/internal/domain/models"
)

type MySqlClosableTransactionable struct {
	connection *sql.Conn
	ctx        context.Context
}

func NewMySqlClosable(ctx context.Context, connection *sql.Conn) *MySqlClosableTransactionable {
	return &MySqlClosableTransactionable{
		connection: connection,
		ctx:        ctx,
	}
}

func (mst *MySqlClosableTransactionable) CloseConnection() *models.CustomError {
	err := mst.connection.Close()

	if err != nil {
		return customerrors.MakeInternalServerError(err.Error(), err)
	}

	return nil
}

func (mst MySqlClosableTransactionable) BeginTransaction() (repository.Transaction, *models.CustomError) {
	tx, err := mst.connection.BeginTx(mst.ctx, nil)
	if err != nil {
		log.Println(err)
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

func (mst *MySqlTransaction) Query(args ...any) (any, *models.CustomError) {
	rawSttmt := args[0]

	switch sttmt := rawSttmt.(type) {
	case string:
		sqlArgs := args[1:]

		result, err := mst.tx.Query(sttmt, sqlArgs...)
		if err != nil {
			return nil, &models.CustomError{
				Code:    http.StatusInternalServerError,
				Message: "Failed on executing transaction",
				Source:  err,
			}
		}

		return result, nil
	default:
		return nil, &models.CustomError{
			Code:    http.StatusBadRequest,
			Message: "Invalid type for query. Must be a string",
		}
	}
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

		return result, nil
	default:
		return nil, &models.CustomError{
			Code:    http.StatusBadRequest,
			Message: "Invalid type for query. Must be a string",
		}
	}
}
