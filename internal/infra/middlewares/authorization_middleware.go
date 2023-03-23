package middlewares

import (
	"net/http"
	"strings"

	"github.com/megalypse/golang-verifymy-backend-test/internal/domain/models"
	"github.com/megalypse/golang-verifymy-backend-test/internal/presentation/http/controllers"
)

func AuthorizationMiddleware(userRoles string, requiredRoles []string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		isAuthorized := true

		for _, role := range requiredRoles {
			isAuthorized = isAuthorized && strings.Contains(userRoles, role)
		}

		if isAuthorized {
			next.ServeHTTP(w, r)
		} else {
			controllers.WriteError(w, &models.CustomError{
				Code:    http.StatusForbidden,
				Message: http.StatusText(http.StatusForbidden),
			})
		}
	})
}
