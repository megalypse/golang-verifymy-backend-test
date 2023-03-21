package factory

import (
	"github.com/megalypse/golang-verifymy-backend-test/internal/data/repository"
	"github.com/megalypse/golang-verifymy-backend-test/internal/infra/db/repositorymysql"
)

var addressRepository repository.AddressRepository

func init() {
	addressRepository = repositorymysql.MySqlAddressRepository{}
}

func GetMySqlAddressRepository() repository.AddressRepository {
	return addressRepository
}
