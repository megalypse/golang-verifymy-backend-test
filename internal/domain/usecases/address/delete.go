package address

import (
	"context"

	"github.com/megalypse/golang-verifymy-backend-test/internal/domain/models"
)

type DeleteAddress interface {
	Delete(ctx context.Context, id int64) *models.CustomError
}
