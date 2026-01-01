package auth

import (
	"app/adv-http/configs"
	"fmt"
	"net/http"
)

type AuthHendlerDeps struct {
	*configs.Config
}
type AuthHandler struct {
	*configs.Config
}

func NewAuthHandler(router *http.ServeMux, deps *AuthHendlerDeps) {
	handler := &AuthHandler{
		Config: deps.Config,
	}
	router.HandleFunc("POST /auth/login", handler.Login())
	router.HandleFunc("POST /auth/register", handler.Register())
}

func (handler *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		secret := handler.Config.Auth.Secret
		fmt.Fprintf(w, "Login! %s", secret)
	}
}

func (handler *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Register!")
	}
}
