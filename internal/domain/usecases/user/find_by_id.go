package usecases

import "github.com/megalypse/golang-verifymy-backend-test/internal/domain/models"

type FindUserById interface {
	FindById(int64) (*models.User, *models.CustomError)
}
