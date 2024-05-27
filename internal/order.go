package domain

import (
	"time"

	"github.com/google/uuid"
)

type orderDomain struct {
	id        uuid.UUID
	amount    float64
	payee     uuid.UUID
	payer     uuid.UUID
	createdAt time.Time
}

type OrderDomainInterface interface {
	GetID() uuid.UUID
	GetAmount() float64
	GetPayee() uuid.UUID
	GetPayer() uuid.UUID
	GetCreatedAt() time.Time
}

func NewOrderDomain(
	amount float64,
	payee uuid.UUID,
	payer uuid.UUID,
) *orderDomain {
	return &orderDomain{
		id:        uuid.New(),
		amount:    amount,
		payee:     payee,
		payer:     payer,
		createdAt: time.Now(),
	}
}

func NewOrderUpdateDomain(
	amount float64,
	payee uuid.UUID,
	payer uuid.UUID,
) OrderDomainInterface {
	return &orderDomain{
		amount:    amount,
		payee:     payee,
		payer:     payer,
		createdAt: time.Now(),
	}
}

func (o *orderDomain) GetID() uuid.UUID {
	return o.id
}

func (o *orderDomain) GetAmount() float64 {
	return o.amount
}

func (o *orderDomain) SetAmount(amount float64) {
	o.amount = amount
}

func (o *orderDomain) GetPayee() uuid.UUID {
	return o.payee
}

func (o *orderDomain) SetPayee(payee uuid.UUID) {
	o.payee = payee
}

func (o *orderDomain) GetPayer() uuid.UUID {
	return o.payer
}

func (o *orderDomain) GetCreatedAt() time.Time {
	return o.createdAt
}
