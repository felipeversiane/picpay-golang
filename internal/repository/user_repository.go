package repository

import (
	"context"

	domain "github.com/felipeversiane/picpay-golang.git/internal"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type userRepository struct {
	Conn *pgxpool.Pool
}

type UserRepository interface {
	InsertUserRepository(ctx context.Context, user domain.User) (domain.User, error)
	FindUserByIDRepository(ctx context.Context, id uuid.UUID) (domain.User, error)
	FindUserByEmailRepository(ctx context.Context, email string) (domain.User, error)
	UpdateUserRepository(ctx context.Context, user domain.User) error
	DeleteUserRepository(ctx context.Context, id uuid.UUID) error
}

func (ur *userRepository) InsertUserRepository(ctx context.Context, user domain.User) (domain.User, error) {
	return domain.User{}, nil
}

func (ur *userRepository) FindUserByIDRepository(ctx context.Context, id uuid.UUID) (domain.User, error) {
	return domain.User{}, nil
}

func (ur *userRepository) FindUserByEmailRepository(ctx context.Context, email string) (domain.User, error) {
	return domain.User{}, nil
}

func (ur *userRepository) UpdateUserRepository(ctx context.Context, user domain.User) error {
	return nil
}

func (ur *userRepository) DeleteUserRepository(ctx context.Context, id uuid.UUID) error {
	return nil
}
