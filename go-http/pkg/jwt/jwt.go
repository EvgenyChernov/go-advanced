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
		// Проверяем метод подписи (должен быть HMAC)
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(j.Secret), nil
	})
	if err != nil || token == nil {
		return false, nil
	}
	
	// Безопасная проверка типа claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return false, nil
	}
	
	// Безопасное извлечение email
	email, ok := claims["email"].(string)
	if !ok || email == "" {
		return false, nil
	}
	
	return true, &JWTData{Email: email}
}
