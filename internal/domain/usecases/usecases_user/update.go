package usecases_user

import "github.com/megalypse/golang-verifymy-backend-test/internal/domain/models"

type UpdateUser interface {
	Update(models.User) *models.User
}
