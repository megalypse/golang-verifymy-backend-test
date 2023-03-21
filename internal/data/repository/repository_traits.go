package repository

import "github.com/megalypse/golang-verifymy-backend-test/internal/domain/models"

type Connectable interface {
	BeginTransaction() (Transaction, *models.CustomError)
	CloseConnection()
}

type Transaction interface {
	Rollback() *models.CustomError
	Commit() *models.CustomError
	Exec(...any) (any, *models.CustomError)
}
