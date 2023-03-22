package userrepository

import (
	mysqlrepository "github.com/megalypse/golang-verifymy-backend-test/internal/infra/repository/mysql"
)

type MySqlUserRepository struct {
	mysqlrepository.BaseMySqlRepository
}
