package middleware

import (
	"github.com/hertz-contrib/migrate/cmd/net/testdata/GoWeb/security"
	"log/slog"
	"net/http"
)

// Csrf validates the CSRF token and returns the handler function if it succeeded
func Csrf(f func(w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := security.VerifyCsrfToken(r)
		if err != nil {
			slog.Info("error verifying csrf token")
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}

		f(w, r)
	}
}
