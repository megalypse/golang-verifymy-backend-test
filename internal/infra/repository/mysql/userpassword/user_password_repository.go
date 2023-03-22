package userpasswordrepository

import (
	mysqlrepository "github.com/megalypse/golang-verifymy-backend-test/internal/infra/repository/mysql"
)

type MySqlUserPasswordRepository struct {
	mysqlrepository.BaseMySqlRepository
}
