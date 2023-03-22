package repositorymysql

import (
	"context"

	"github.com/megalypse/golang-verifymy-backend-test/internal/data/repository"
	"github.com/megalypse/golang-verifymy-backend-test/internal/infra/repository/mysql/config"
)

type BaseMySqlRepository struct{}

func (BaseMySqlRepository) NewConnection(ctx context.Context) repository.Closable {
	newConnection := config.GetMySqlConnection(ctx)

	return NewMySqlClosable(ctx, newConnection)
}
