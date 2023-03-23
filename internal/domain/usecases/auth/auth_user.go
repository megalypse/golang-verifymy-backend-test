package auth

import (
	"context"

	"github.com/megalypse/golang-verifymy-backend-test/internal/domain/models"
)

type AuthUser interface {
	Auth(context.Context, *models.User) (bool, *models.CustomError)
}
