package middleware

import (
	"app/adv-http/configs"
	"app/adv-http/pkg/jwt"
	"context"
	"net/http"
	"strings"
)

type key string

const (
	ContextKeyEmail key = "ContextKeyEmail"
)

func writeUnauthorized(w http.ResponseWriter) {
	w.WriteHeader(http.StatusUnauthorized)
	w.Write([]byte(http.StatusText(http.StatusUnauthorized)))
}

func IsAuthenticated(next http.Handler, config *configs.Config) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if !strings.HasPrefix(authHeader, "Baerer ") {
			writeUnauthorized(w)
			return
		}
		token := strings.TrimPrefix(authHeader, "Bearer ")
		isValid, data := jwt.NewJWT(config.Auth.Secret).Parse(token)
		if !isValid {
			writeUnauthorized(w)
			return
		}
		ctx := context.WithValue(r.Context(), ContextKeyEmail, data.Email)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
