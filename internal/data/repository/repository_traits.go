package repository

import "github.com/megalypse/golang-verifymy-backend-test/internal/domain/models"

type Closable interface {
	BeginTransaction() (Transaction, *models.CustomError)
	CloseConnection()
}

type Transaction interface {
	Rollback() *models.CustomError
	Commit() *models.CustomError

	// Use this one for operations that do not return rows, operations like INSERT or UPDATE
	Exec(...any) (any, *models.CustomError)

	// Use this one for operations that do return rows, operations like SELECT
	Query(...any) (any, *models.CustomError)
}
