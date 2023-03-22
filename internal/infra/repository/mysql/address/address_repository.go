package addressrepository

import (
	mysqlrepository "github.com/megalypse/golang-verifymy-backend-test/internal/infra/repository/mysql"
)

type MySqlAddressRepository struct {
	mysqlrepository.BaseMySqlRepository
}
