package factory

import (
	"context"

	"github.com/megalypse/golang-verifymy-backend-test/internal/data/repository"
	repositorymysql "github.com/megalypse/golang-verifymy-backend-test/internal/infra/repository/mysql"
	"github.com/megalypse/golang-verifymy-backend-test/internal/infra/repository/mysql/config"
)

func NewSqlConnection(ctx context.Context) repository.Closable {
	newConnection := config.GetMySqlConnection(ctx)

	return repositorymysql.NewMySqlClosable(ctx, newConnection)
}
