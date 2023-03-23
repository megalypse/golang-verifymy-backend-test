package security

import (
	"context"

	"github.com/megalypse/golang-verifymy-backend-test/internal/domain/models"
)

type AuthorizationService interface {
	AssignRole(ctx context.Context, userId int64, roleAlias string) *models.CustomError
}
