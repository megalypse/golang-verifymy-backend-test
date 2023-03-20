package usecases_user

import "github.com/megalypse/golang-verifymy-backend-test/internal/domain/models"

type DeleteUser interface {
	Delete(models.User) bool
}
