package factory

import (
	"github.com/megalypse/golang-verifymy-backend-test/internal/data/repository"
	repositorymysql "github.com/megalypse/golang-verifymy-backend-test/internal/infra/repository/mysql"
)

var addressRepository repository.AddressRepository

func init() {
	addressRepository = repositorymysql.MySqlAddressRepository{}
}

func GetMySqlAddressRepository() repository.AddressRepository {
	return addressRepository
}
