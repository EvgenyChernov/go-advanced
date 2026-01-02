package auth

import (
	"app/adv-http/configs"
	"app/adv-http/pkg/response"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
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

		var payload LoginRequest

		err := json.NewDecoder(req.Body).Decode(&payload)
		if err != nil {
			response.JsonResponse(w, err.Error(), 402)
			return
		}
		validate := validator.New()
		err = validate.Struct(payload)
		if err != nil {
			response.JsonResponse(w, err.Error(), 402)
			return
		}
		fmt.Println(payload)

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
