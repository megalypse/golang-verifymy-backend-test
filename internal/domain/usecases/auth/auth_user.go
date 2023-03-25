package auth

import (
	"context"

	"github.com/megalypse/golang-verifymy-backend-test/internal/domain/models"
)

type UserSignIn interface {
	SignIn(context.Context, *models.User) (*models.User, *models.CustomError)
}
