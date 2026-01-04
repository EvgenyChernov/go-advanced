package auth

import (
	"app/adv-http/internal/user"
)

type AuthService struct {
	UserRepository *user.UserRepository
}

func NewAuthService(userRepository *user.UserRepository) *AuthService {
	return &AuthService{
		UserRepository: userRepository,
	}
}

func (service *AuthService) Register(email string, password string, name string) (string, error) {
	existingUser, _ := service.UserRepository.FindByEmail(email)
	if existingUser != nil {
		return "", ErrUserAlreadyExists
	}
	user := &user.User{
		Email:    email,
		Password: "",
		Name:     name,
	}
	_, err := service.UserRepository.Create(user)
	if err != nil {
		return "", err
	}
	return user.Email, nil
}
