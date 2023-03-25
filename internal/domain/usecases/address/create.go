package address

import (
	"context"

	"github.com/megalypse/golang-verifymy-backend-test/internal/domain/models"
)

type CreateAddress interface {
	Create(ctx context.Context, source *models.Address) (*models.Address, *models.CustomError)
}
