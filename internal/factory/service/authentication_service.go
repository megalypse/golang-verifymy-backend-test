package service

import (
	"github.com/megalypse/golang-verifymy-backend-test/internal/data/services/security"
	"github.com/megalypse/golang-verifymy-backend-test/internal/infra/service"
)

var authenticationService security.AuthenticationService

func init() {
	authenticationService = service.JwtAuthenticationService{}
}

func GetJwtAuthenticationService() security.AuthenticationService {
	return authenticationService
}
