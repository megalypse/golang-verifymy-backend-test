package service

import (
	"github.com/megalypse/golang-verifymy-backend-test/internal/data/repository"
	"github.com/megalypse/golang-verifymy-backend-test/internal/data/services/security"
)

type UserService struct {
	userRepository         repository.UserRepository
	userPasswordRepository repository.UserPasswordRepository
	addressRepository      repository.AddressRepository
	securityService        security.SecurityService
}

func NewUserService(
	userRepository repository.UserRepository,
	userPasswordRepository repository.UserPasswordRepository,
	addressRepository repository.AddressRepository,
	securityService security.SecurityService,
) UserService {
	return UserService{
		userRepository:         userRepository,
		userPasswordRepository: userPasswordRepository,
		addressRepository:      addressRepository,
		securityService:        securityService,
	}
}
