package middle

import (
	"net/http"
)

func CheckRole(requiredRole string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			role, ok := r.Context().Value("role").(string)
			if !ok || role == "" {
				http.Error(w, "Role not found", http.StatusUnauthorized)
				return
			}

			if role != requiredRole {
				http.Error(w, "Access Denied", http.StatusForbidden)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
