package repository

import (
	"context"

	"github.com/felipeversiane/picpay-golang.git/config/http_error"
	domain "github.com/felipeversiane/picpay-golang.git/internal"
	"github.com/felipeversiane/picpay-golang.git/internal/entity/response"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type userRepository struct {
	conn *pgxpool.Pool
}

func NewUserRepository(
	conn *pgxpool.Pool,
) UserRepository {
	return &userRepository{
		conn,
	}
}

type UserRepository interface {
	InsertUserRepository(ctx context.Context, user domain.UserDomainInterface) (response.UserResponse, *http_error.HttpError)
	FindUserByDocumentRepository(ctx context.Context, document string) (response.UserResponse, *http_error.HttpError)
	FindUserByIDRepository(ctx context.Context, id uuid.UUID) (response.UserResponse, *http_error.HttpError)
	FindUserByEmailRepository(ctx context.Context, email string) (response.UserResponse, *http_error.HttpError)
	UpdateUserRepository(ctx context.Context, user domain.UserDomainInterface, id uuid.UUID) (response.UserResponse, *http_error.HttpError)
	DeleteUserRepository(ctx context.Context, id uuid.UUID) *http_error.HttpError
}

func (ur *userRepository) InsertUserRepository(ctx context.Context, user domain.UserDomainInterface) (response.UserResponse, *http_error.HttpError) {
	query := "INSERT INTO users (id, email, password, first_name, last_name, document, balance, is_merchant, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id, email, password, first_name, last_name, document, balance, is_merchant, created_at, updated_at"
	var insertedUser response.UserResponse
	err := ur.conn.QueryRow(ctx, query,
		user.GetID(), user.GetEmail(),
		user.GetPassword(), user.GetFirstName(),
		user.GetLastName(), user.GetDocument(),
		user.GetBalance(), user.GetIsMerchant(),
		user.GetCreatedAt(), user.GetUpdatedAt()).Scan(
		&insertedUser.ID, &insertedUser.Email,
		&insertedUser.Password, &insertedUser.FirstName,
		&insertedUser.LastName, &insertedUser.Document,
		&insertedUser.Balance, &insertedUser.IsMerchant,
		&insertedUser.CreatedAt, &insertedUser.UpdatedAt,
	)

	if err != nil {
		return response.UserResponse{}, http_error.NewInternalServerError(err.Error())
	}

	return insertedUser, nil
}

func (ur *userRepository) FindUserByDocumentRepository(ctx context.Context, document string) (response.UserResponse, *http_error.HttpError) {
	query := "SELECT id, email, password, first_name, last_name, document, balance, is_merchant, created_at, updated_at FROM users WHERE document = $1"
	var foundUser response.UserResponse
	err := ur.conn.QueryRow(ctx, query, document).Scan(
		&foundUser.ID, &foundUser.Email,
		&foundUser.Password, &foundUser.FirstName,
		&foundUser.LastName, &foundUser.Document,
		&foundUser.Balance, &foundUser.IsMerchant,
		&foundUser.CreatedAt, &foundUser.UpdatedAt,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return response.UserResponse{}, http_error.NewNotFoundError("User not found")
		}
		return response.UserResponse{}, http_error.NewInternalServerError(err.Error())
	}

	return foundUser, nil
}

func (ur *userRepository) FindUserByIDRepository(ctx context.Context, id uuid.UUID) (response.UserResponse, *http_error.HttpError) {
	query := "SELECT id, email, password, first_name, last_name, document, balance, is_merchant, created_at, updated_at FROM users WHERE id = $1"
	var foundUser response.UserResponse
	err := ur.conn.QueryRow(ctx, query, id).Scan(
		&foundUser.ID, &foundUser.Email,
		&foundUser.Password, &foundUser.FirstName,
		&foundUser.LastName, &foundUser.Document,
		&foundUser.Balance, &foundUser.IsMerchant,
		&foundUser.CreatedAt, &foundUser.UpdatedAt,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return response.UserResponse{}, http_error.NewNotFoundError("User not found")
		}
		return response.UserResponse{}, http_error.NewInternalServerError(err.Error())
	}

	return foundUser, nil
}

func (ur *userRepository) FindUserByEmailRepository(ctx context.Context, email string) (response.UserResponse, *http_error.HttpError) {
	query := "SELECT id, email, password, first_name, last_name, document, balance, is_merchant, created_at, updated_at FROM users WHERE email = $1"
	var foundUser response.UserResponse
	err := ur.conn.QueryRow(ctx, query, email).Scan(
		&foundUser.ID, &foundUser.Email,
		&foundUser.Password, &foundUser.FirstName,
		&foundUser.LastName, &foundUser.Document,
		&foundUser.Balance, &foundUser.IsMerchant,
		&foundUser.CreatedAt, &foundUser.UpdatedAt,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return response.UserResponse{}, http_error.NewNotFoundError("User not found")
		}
		return response.UserResponse{}, http_error.NewInternalServerError(err.Error())
	}

	return foundUser, nil
}

func (ur *userRepository) UpdateUserRepository(ctx context.Context, user domain.UserDomainInterface, id uuid.UUID) (response.UserResponse, *http_error.HttpError) {
	query := `
		UPDATE users 
		SET 
			first_name = $1, 
			last_name = $2, 
			balance = $3, 
			is_merchant = $4, 
			updated_at = now() 
		WHERE 
			id = $5
		RETURNING id, email, password, first_name, last_name, document, balance, is_merchant, created_at, updated_at
	`

	var updatedUser response.UserResponse
	err := ur.conn.QueryRow(ctx, query,
		user.GetFirstName(),
		user.GetLastName(),
		user.GetBalance(),
		user.GetIsMerchant(),
		id,
	).Scan(
		&updatedUser.ID,
		&updatedUser.Email,
		&updatedUser.Password,
		&updatedUser.FirstName,
		&updatedUser.LastName,
		&updatedUser.Document,
		&updatedUser.Balance,
		&updatedUser.IsMerchant,
		&updatedUser.CreatedAt,
		&updatedUser.UpdatedAt,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return response.UserResponse{}, http_error.NewNotFoundError("User not found")
		}
		return response.UserResponse{}, http_error.NewInternalServerError(err.Error())
	}

	return updatedUser, nil
}

func (ur *userRepository) DeleteUserRepository(ctx context.Context, id uuid.UUID) *http_error.HttpError {
	query := "DELETE FROM users WHERE id = $1"
	_, err := ur.conn.Exec(ctx, query, id)
	if err != nil {
		return http_error.NewInternalServerError(err.Error())
	}
	return nil
}
