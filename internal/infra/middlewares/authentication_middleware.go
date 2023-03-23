package middlewares

import (
	"context"
	"net/http"
	"strings"
	"time"

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
			claims, ok := jwtToken.Claims.(jwt.MapClaims)
			if !ok {
				controllers.WriteError(w, customerrors.MakeInternalServerError("Failed on getting token claims", nil))
				return
			}

			isValid := checkTokenExpiration(claims)
			if !isValid {
				controllers.WriteError(w, unauthorizedError)
				return
			}

			roles := claims["roles"].(string)
			ctx := context.WithValue(r.Context(), "roles", roles)

			next.ServeHTTP(w, r.WithContext(ctx))
		} else {
			controllers.WriteError(w, unauthorizedError)
			return
		}

	})
}

func checkTokenExpiration(claims jwt.MapClaims) bool {
	rawExpiresAt := claims["expires_at"].(string)
	layout := "2006-01-02T15:04:05.99999999Z"
	expiresAt, err := time.Parse(layout, rawExpiresAt)
	if err != nil {
		panic(err)
	}

	if time.Now().Unix() > expiresAt.Unix() {
		return false
	}

	return true
}
