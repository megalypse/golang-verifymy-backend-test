package service

import (
	"github.com/megalypse/golang-verifymy-backend-test/internal/data/services/security"
	repositoryFactory "github.com/megalypse/golang-verifymy-backend-test/internal/factory/repository/mysql"
	"github.com/megalypse/golang-verifymy-backend-test/internal/infra/service"
)

var authorizationService security.AuthorizationService

func init() {
	authorizationService = service.RolesAuthorizationService{
		RolesRepository:     repositoryFactory.GetRolesRepository(),
		UserRolesRepository: repositoryFactory.GetUserRolesRepository(),
	}
}

func GetRolesAuthorizationService() security.AuthorizationService {
	return authorizationService
}
