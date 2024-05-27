package request

type OrderRequest struct {
	Amount float64 `json:"amount" binding:"required,numeric,min=0.01"`
	Payee  string  `json:"payee" binding:"required"`
	Payer  string  `json:"payer" binding:"required"`
}
