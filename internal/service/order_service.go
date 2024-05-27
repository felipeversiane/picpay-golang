package service

import (
	"github.com/felipeversiane/picpay-golang.git/internal/repository"
)

type orderService struct {
	orderRepository repository.OrderRepository
}

func NewOrderRepository(
	orderRepository repository.OrderRepository,
) OrderService {
	return &orderService{
		orderRepository,
	}
}

type OrderService interface {
}
