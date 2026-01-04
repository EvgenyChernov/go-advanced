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
		// Проверяем наличие заголовка Authorization с префиксом "Bearer "
		if !strings.HasPrefix(authHeader, "Bearer ") {
			writeUnauthorized(w)
			return
		}
		// Извлекаем токен (удаляем префикс "Bearer ")
		token := strings.TrimPrefix(authHeader, "Bearer ")
		if token == "" {
			writeUnauthorized(w)
			return
		}
		// Парсим и проверяем токен
		isValid, data := jwt.NewJWT(config.Auth.Secret).Parse(token)
		if !isValid || data == nil {
			writeUnauthorized(w)
			return
		}
		// Добавляем email в контекст для использования в handler
		ctx := context.WithValue(r.Context(), ContextKeyEmail, data.Email)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
