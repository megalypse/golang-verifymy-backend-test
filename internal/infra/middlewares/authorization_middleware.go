package middlewares

import (
	"net/http"
	"strings"

	"github.com/megalypse/golang-verifymy-backend-test/internal/domain/models"
	httputils "github.com/megalypse/golang-verifymy-backend-test/internal/presentation/http"
)

func AuthorizationMiddleware(requiredRoles []string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		isAuthorized := true
		userRoles := r.Context().Value("roles").(string)

		for _, role := range requiredRoles {
			isAuthorized = isAuthorized && strings.Contains(userRoles, role)
		}

		if isAuthorized {
			next.ServeHTTP(w, r)
		} else {
			httputils.WriteError(w, &models.CustomError{
				Code:    http.StatusForbidden,
				Message: http.StatusText(http.StatusForbidden),
			})
		}
	})
}
