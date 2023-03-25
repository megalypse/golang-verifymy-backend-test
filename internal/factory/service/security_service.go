package service

import (
	"github.com/megalypse/golang-verifymy-backend-test/internal/data/services/security"
	"github.com/megalypse/golang-verifymy-backend-test/internal/infra/service"
)

var securityService security.SecurityService

func init() {
	securityService = service.BCryptSecurityService{}
}

func GetBCryptSecurityService() security.SecurityService {
	return securityService
}
