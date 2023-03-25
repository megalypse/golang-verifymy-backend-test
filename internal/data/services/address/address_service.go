package address

import (
	"github.com/megalypse/golang-verifymy-backend-test/internal/data/repository"
)

type AddressService struct {
	AddressRepository repository.AddressRepository
	UserRepository    repository.UserRepository
}
