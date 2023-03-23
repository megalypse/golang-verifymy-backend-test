package middlewares

import (
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/megalypse/golang-verifymy-backend-test/config"
	"github.com/megalypse/golang-verifymy-backend-test/internal/domain/customerrors"
	"github.com/megalypse/golang-verifymy-backend-test/internal/domain/models"
	"github.com/megalypse/golang-verifymy-backend-test/internal/presentation/http/controllers"
)

func VerifyJwt(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		secret := config.GetAuthSecret()
		unauthorizedError := &models.CustomError{
			Code:    http.StatusUnauthorized,
			Message: "Authentication is required",
		}

		rawToken := r.Header.Get("Authorization")
		if rawToken == "" {
			controllers.WriteError(w, unauthorizedError)
			return
		}

		token := strings.TrimPrefix(rawToken, "Bearer")
		token = strings.TrimSpace(token)
		jwtToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, customerrors.MakeInternalServerError("Failed parsing json token", nil)
			}

			return []byte(secret), nil
		})
		if err != nil {
			controllers.WriteError(w, &models.CustomError{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
				Source:  err,
			})
			return
		}

		if jwtToken.Valid {
			next.ServeHTTP(w, r)
		} else {
			controllers.WriteError(w, unauthorizedError)
			return
		}

	})
}
