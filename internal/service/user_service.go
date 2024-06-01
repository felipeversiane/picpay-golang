package service

import (
	"context"

	"github.com/felipeversiane/picpay-golang.git/config/http_error"
	"github.com/felipeversiane/picpay-golang.git/config/logger"
	domain "github.com/felipeversiane/picpay-golang.git/internal"
	"github.com/felipeversiane/picpay-golang.git/internal/entity/response"
	"github.com/felipeversiane/picpay-golang.git/internal/repository"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(
	userRepository repository.UserRepository,
) UserService {
	return &userService{
		userRepository,
	}
}

type UserService interface {
	InsertUserService(ctx context.Context, user domain.UserDomainInterface) (
		response.UserResponse, *http_error.HttpError)
	FindUserByDocumentService(
		document string, ctx context.Context,
	) (response.UserResponse, *http_error.HttpError)
	FindUserByIDService(
		id uuid.UUID, ctx context.Context,
	) (response.UserResponse, *http_error.HttpError)
	FindUserByEmailService(
		email string, ctx context.Context,
	) (response.UserResponse, *http_error.HttpError)
	UpdateUserService(id uuid.UUID, user domain.UserDomainInterface, ctx context.Context) (response.UserResponse, *http_error.HttpError)
	DeleteUserService(id uuid.UUID, ctx context.Context) *http_error.HttpError
}

func (uc *userService) InsertUserService(ctx context.Context, user domain.UserDomainInterface) (response.UserResponse, *http_error.HttpError) {
	user.EncryptPassword()
	_, err := uc.FindUserByDocumentService(user.GetDocument(), ctx)
	if err == nil {
		return response.UserResponse{}, http_error.NewBadRequestError("Document is already registered in another account")
	}
	_, err = uc.FindUserByEmailService(user.GetEmail(), ctx)
	if err == nil {
		return response.UserResponse{}, http_error.NewBadRequestError("Email is already registered in another account")
	}
	result, err := uc.userRepository.InsertUserRepository(ctx, user)
	if err != nil {
		logger.Error("Error trying to call repository",
			err,
			zap.String("journey", "InsertUser"))
		return response.UserResponse{}, err
	}
	return result, nil
}

func (uc *userService) FindUserByDocumentService(document string, ctx context.Context) (response.UserResponse, *http_error.HttpError) {
	result, err := uc.userRepository.FindUserByDocumentRepository(ctx, document)
	if err != nil {
		logger.Error("Error trying to call repository",
			err,
			zap.String("journey", "FindUserByDocument"))
		return response.UserResponse{}, err
	}
	return result, nil
}

func (uc *userService) FindUserByIDService(id uuid.UUID, ctx context.Context) (response.UserResponse, *http_error.HttpError) {
	result, err := uc.userRepository.FindUserByIDRepository(ctx, id)
	if err != nil {
		logger.Error("Error trying to call repository",
			err,
			zap.String("journey", "FindUserByDocument"))
		return response.UserResponse{}, err
	}
	return result, nil
}

func (uc *userService) FindUserByEmailService(email string, ctx context.Context) (response.UserResponse, *http_error.HttpError) {
	result, err := uc.userRepository.FindUserByEmailRepository(ctx, email)
	if err != nil {
		logger.Error("Error trying to call repository",
			err,
			zap.String("journey", "FindUserByEmail"))
		return response.UserResponse{}, err
	}
	return result, nil
}

func (uc *userService) UpdateUserService(id uuid.UUID, user domain.UserDomainInterface, ctx context.Context) (response.UserResponse, *http_error.HttpError) {
	_, err := uc.FindUserByIDService(id, ctx)
	if err != nil {
		return response.UserResponse{}, http_error.NewNotFoundError("User not found")
	}

	result, err := uc.userRepository.UpdateUserRepository(ctx, user, id)
	if err != nil {
		logger.Error("Error trying to call repository",
			err,
			zap.String("journey", "UpdateUser"))
		return response.UserResponse{}, err
	}
	return result, nil
}

func (uc *userService) DeleteUserService(id uuid.UUID, ctx context.Context) *http_error.HttpError {
	_, err := uc.FindUserByIDService(id, ctx)
	if err != nil {
		return http_error.NewNotFoundError("User not found")
	}

	err = uc.userRepository.DeleteUserRepository(ctx, id)
	if err != nil {
		logger.Error("Error trying to call repository",
			err,
			zap.String("journey", "DeleteUser"))
		return err
	}
	return nil
}
