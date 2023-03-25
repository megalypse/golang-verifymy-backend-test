package address

import (
	"context"

	"github.com/megalypse/golang-verifymy-backend-test/internal/domain/models"
)

type UpdateAddress interface {
	Update(context.Context, *models.Address) (*models.Address, *models.CustomError)
}
