package security

import (
	"context"

	"github.com/megalypse/golang-verifymy-backend-test/internal/domain/models"
)

type AuthorizationService interface {
	AssignRole(ctx context.Context, userId int64, roleId int64) *models.CustomError
	GetUserRoles(ctx context.Context, userId int64) ([]models.Role, *models.CustomError)
}
