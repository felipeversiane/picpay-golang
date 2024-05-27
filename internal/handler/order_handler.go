package handler

import (
	"github.com/felipeversiane/picpay-golang.git/internal/service"
)

type orderHandler struct {
	orderService service.OrderService
}

func NewOrderHandler(
	orderService service.OrderService,
) OrderHandler {
	return &orderHandler{
		orderService,
	}
}

type OrderHandler interface {
}
