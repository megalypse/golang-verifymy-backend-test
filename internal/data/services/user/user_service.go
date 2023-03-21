package service

import (
	"github.com/megalypse/golang-verifymy-backend-test/internal/data/repository"
)

type UserService struct {
	userRepository         repository.UserRepository
	userPasswordRepository repository.UserPasswordRepository
	addressRepository      repository.AddressRepository
}

func NewUserService(
	userRepository repository.UserRepository,
	userPasswordRepository repository.UserPasswordRepository,
	addressRepository repository.AddressRepository,
) UserService {
	return UserService{
		userRepository:         userRepository,
		userPasswordRepository: userPasswordRepository,
		addressRepository:      addressRepository,
	}
}
