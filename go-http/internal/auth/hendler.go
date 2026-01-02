package auth

import (
	"app/adv-http/configs"
	"app/adv-http/pkg/response"
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
		fmt.Println(secret)
		data := LoginResponse{
			Token: secret,
		}
		response.JsonResponse(w, data, 200)
	}
}

func (handler *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Register!")
	}
}
