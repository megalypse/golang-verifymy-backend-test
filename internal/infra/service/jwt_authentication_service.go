package service

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/megalypse/golang-verifymy-backend-test/config"
	"github.com/megalypse/golang-verifymy-backend-test/internal/domain/customerrors"
	"github.com/megalypse/golang-verifymy-backend-test/internal/domain/models"
)

type JwtAuthenticationService struct{}

func (JwtAuthenticationService) GenerateJwtToken(userEmail string, roles []string) (string, *models.CustomError) {
	rawSecret := config.GetAuthSecret()
	if rawSecret == "" {
		panic("Auth secret not defined")
	}

	secret := []byte(rawSecret)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"expires_at": time.Now().Add(time.Hour),
		"email":      userEmail,
		"roles":      stringifyRoles(roles),
	})

	signedToken, err := token.SignedString(secret)
	if err != nil {
		return "", customerrors.MakeInternalServerError(err.Error(), nil)
	}

	return signedToken, nil
}

func stringifyRoles(roles []string) string {
	var rolesString string

	for i, v := range roles {
		if i != 0 {
			rolesString += ","
		}

		rolesString += v
	}

	return rolesString
}
