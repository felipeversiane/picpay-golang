package service

import (
	"github.com/felipeversiane/picpay-golang.git/config/http_error"
	domain "github.com/felipeversiane/picpay-golang.git/internal"
	"github.com/felipeversiane/picpay-golang.git/internal/repository"
	"github.com/google/uuid"
)

type userService struct {
	userRepository repository.UserRepository
}

type UserService interface {
	InsertUserService(domain.User) (
		domain.User, *http_error.HttpError)

	FindUserByIDService(
		id uuid.UUID,
	) (domain.User, *http_error.HttpError)
	FindUserByEmailService(
		email string,
	) (domain.User, *http_error.HttpError)

	UpdateUserService(uuid.UUID, domain.User) *http_error.HttpError
	DeleteUserService(uuid.UUID) *http_error.HttpError
}

func (uc *userService) InsertUserService(user domain.User) (domain.User, *http_error.HttpError) {
	return domain.User{}, nil
}

func (uc *userService) FindUserByIDService(id uuid.UUID) (domain.User, *http_error.HttpError) {
	return domain.User{}, nil
}

func (uc *userService) FindUserByEmailService(email string) (domain.User, *http_error.HttpError) {
	return domain.User{}, nil
}

func (uc *userService) UpdateUserService(id uuid.UUID, user domain.User) *http_error.HttpError {
	return nil
}

func (uc *userService) DeleteUserService(id uuid.UUID) *http_error.HttpError {
	return nil
}
