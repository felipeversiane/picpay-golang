package repository

import (
	"context"
	"time"

	"github.com/felipeversiane/picpay-golang.git/config/http_error"
	domain "github.com/felipeversiane/picpay-golang.git/internal"
	"github.com/felipeversiane/picpay-golang.git/internal/entity/response"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type orderRepository struct {
	conn *pgxpool.Pool
}

func NewOrderRepository(
	conn *pgxpool.Pool,
) OrderRepository {
	return &orderRepository{
		conn,
	}
}

type OrderRepository interface {
	InsertOrderRepository(ctx context.Context, order domain.OrderDomainInterface) (response.OrderResponse, *http_error.HttpError)
	FindOrderByIDRepository(ctx context.Context, orderID uuid.UUID) (response.OrderResponse, *http_error.HttpError)
}

func (r *orderRepository) InsertOrderRepository(ctx context.Context, order domain.OrderDomainInterface) (response.OrderResponse, *http_error.HttpError) {
	query := `
		INSERT INTO orders (id, amount, payee, payer, created_at)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, amount, payee, payer, created_at;
	`

	row := r.conn.QueryRow(ctx, query, order.GetID(), order.GetAmount(), order.GetPayee(), order.GetPayer(), time.Now())

	var orderResponse response.OrderResponse
	err := row.Scan(
		&orderResponse.ID,
		&orderResponse.Amount,
		&orderResponse.Payee,
		&orderResponse.Payer,
		&orderResponse.CreatedAt,
	)

	if err != nil {
		return response.OrderResponse{}, http_error.NewInternalServerError(err.Error())
	}

	return orderResponse, nil
}

func (r *orderRepository) FindOrderByIDRepository(ctx context.Context, orderID uuid.UUID) (response.OrderResponse, *http_error.HttpError) {
	query := `
		SELECT id, amount, payee, payer, created_at
		FROM orders
		WHERE id = $1;
	`

	row := r.conn.QueryRow(ctx, query, orderID)

	var order response.OrderResponse
	err := row.Scan(
		&order.ID,
		&order.Amount,
		&order.Payee,
		&order.Payer,
		&order.CreatedAt,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return response.OrderResponse{}, http_error.NewNotFoundError("Order not found")
		}
		return response.OrderResponse{}, http_error.NewInternalServerError(err.Error())
	}

	return order, nil
}
