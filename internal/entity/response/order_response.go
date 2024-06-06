package response

import (
	"time"

	"github.com/google/uuid"
)

type OrderResponse struct {
	ID         uuid.UUID `json:"id"`
	Amount     float64   `json:"amount"`
	Payee      uuid.UUID `json:"payee"`
	Payer      uuid.UUID `json:"payer"`
	CreatedAt  time.Time `json:"created_at"`
	IsReversed time.Time `json:"is_reversed"`
}
