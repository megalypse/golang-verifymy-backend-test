package factory

import (
	"github.com/megalypse/golang-verifymy-backend-test/internal/data/repository"
	mysqlAddressRepository "github.com/megalypse/golang-verifymy-backend-test/internal/infra/repository/mysql/address"
)

var addressRepository repository.AddressRepository

func init() {
	addressRepository = mysqlAddressRepository.MySqlAddressRepository{}
}

func GetAddressRepository() repository.AddressRepository {
	return addressRepository
}
