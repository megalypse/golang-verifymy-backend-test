package repository

import "github.com/megalypse/golang-verifymy-backend-test/internal/domain/models"

type UserRepository interface {
	Create(models.User) *models.User
	Delete(int64) bool
	FindById(int64) *models.User
	Update(models.User) *models.User
}
