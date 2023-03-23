package repository

import "github.com/megalypse/golang-verifymy-backend-test/internal/domain/models"

type UserRolesRepository interface {
	AssignRole(tx Transaction, userId, roleId int64) *models.CustomError
	GetAllByUserId(tx Transaction, userid int64) ([]models.Role, *models.CustomError)
}
