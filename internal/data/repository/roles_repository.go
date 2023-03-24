package repository

import "github.com/megalypse/golang-verifymy-backend-test/internal/domain/models"

type RolesRepository interface {
	FindByAlias(Transaction, string) (*models.Role, *models.CustomError)
}
