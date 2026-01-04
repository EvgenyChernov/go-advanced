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

func IsAuthenticated(next http.Handler, config *configs.Config) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		token := strings.TrimPrefix(authHeader, "Bearer ")
		_, data := jwt.NewJWT(config.Auth.Secret).Parse(token)
		ctx := context.WithValue(r.Context(), ContextKeyEmail, data.Email)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
