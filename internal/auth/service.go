package auth

import (
	"errors"
	"go/links-shorter/internal/user"
)

type AuthService struct {
	UserRepo *user.UserRepository
}

func NewAuthService(userRepo *user.UserRepository) *AuthService {
	return &AuthService{UserRepo: userRepo}
}

func (service *AuthService) Register(email, password, name string) (*user.User, error) {
	_, err := service.UserRepo.GetByEmail(email)

	if err != nil {
		return nil, errors.New(ErrUserExists)
	}

	// TODO: hash password
	user, err := service.UserRepo.Create(user.NewUser(email, name, ""))

	if err != nil {
		return nil, err
	}

	return user, nil
}
