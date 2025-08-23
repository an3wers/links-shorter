package auth

import (
	"errors"
	"go/links-shorter/internal/user"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	UserRepo *user.UserRepository
}

func NewAuthService(userRepo *user.UserRepository) *AuthService {
	return &AuthService{UserRepo: userRepo}
}

func (service *AuthService) Register(email, password, name string) (*user.User, error) {
	_, err := service.UserRepo.GetByEmail(email)

	if err == nil {
		return nil, errors.New(ErrUserExists)
	}

	hashedPasswordBytes, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return nil, err
	}

	user, err := service.UserRepo.Create(user.NewUser(
		email,
		name,
		string(hashedPasswordBytes)),
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (service *AuthService) Login(email, password string) (*user.User, error) {
	user, err := service.UserRepo.GetByEmail(email)

	if err != nil {
		return nil, errors.New(ErrInvalid)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		return nil, errors.New(ErrInvalid)
	}

	return user, nil

}
