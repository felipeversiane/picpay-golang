package repository

import (
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
}
