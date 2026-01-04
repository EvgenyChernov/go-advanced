package jwt

import (
	"github.com/golang-jwt/jwt/v5"
)

type JWT struct {
	Secret string `json:"secret"`
}

type JWTData struct {
	Email string `json:"email"`
}

func NewJWT(secret string) *JWT {
	return &JWT{
		Secret: secret,
	}
}

func (j *JWT) Create(data JWTData) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": data.Email,
	})

	tokenString, err := token.SignedString([]byte(j.Secret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (j *JWT) Parse(tokenString string) (bool, *JWTData) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.Secret), nil
	})
	if err != nil {
		return false, nil
	}
	email := token.Claims.(jwt.MapClaims)["email"]
	return token.Valid, &JWTData{Email: email.(string)}
}
