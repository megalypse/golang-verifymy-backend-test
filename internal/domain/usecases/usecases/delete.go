package usecases

import "github.com/megalypse/golang-verifymy-backend-test/internal/domain/models"

type DeleteUser interface {
	Delete(int64) *models.CustomError
}
