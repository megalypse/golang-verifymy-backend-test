package repository

import "github.com/megalypse/golang-verifymy-backend-test/internal/domain/models"

type UserRepository interface {
	Create(*models.User) (*models.User, *models.CustomError)
	Delete(int64) *models.CustomError
	FindById(int64) (*models.User, *models.CustomError)
	Update(*models.User) (*models.User, *models.CustomError)
}
