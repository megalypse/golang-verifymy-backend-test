package usecases

import "github.com/megalypse/golang-verifymy-backend-test/internal/domain/models"

type CreateUser interface {
	Create(*models.User) (*models.User, *models.CustomError)
}
