package address

import (
	addressService "github.com/megalypse/golang-verifymy-backend-test/internal/data/services/address"
	"github.com/megalypse/golang-verifymy-backend-test/internal/domain/usecases/address"
	factory "github.com/megalypse/golang-verifymy-backend-test/internal/factory/repository/mysql"
)

var createAddressUsecase address.CreateAddress
var updateAddressUsecase address.UpdateAddress
var deleteAddressUsecase address.DeleteAddress

func init() {
	addressService := addressService.AddressService{
		AddressRepository: factory.GetMySqlAddressRepository(),
	}

	createAddressUsecase = addressService
	updateAddressUsecase = addressService
	deleteAddressUsecase = addressService
}

func GetCreateAddressUsecase() address.CreateAddress {
	return createAddressUsecase
}

func GetUpdateAddressUsecase() address.UpdateAddress {
	return updateAddressUsecase
}

func GetDeleteAddressUsecase() address.DeleteAddress {
	return deleteAddressUsecase
}
