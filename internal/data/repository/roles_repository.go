package repository

import "github.com/megalypse/golang-verifymy-backend-test/internal/domain/models"

type RolesRepository interface {
	FindById(Transaction, int64) (*models.Role, *models.CustomError)
}
