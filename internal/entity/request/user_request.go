package request

type UserRequest struct {
	Email      string  `json:"email" binding:"required,email"`
	Password   string  `json:"password" binding:"required,min=6,containsany=!@&*%$#"`
	FirstName  string  `json:"first_name" binding:"required,max=50"`
	LastName   string  `json:"last_name" binding:"required,max=50"`
	Document   string  `json:"document" binding:"required,min=4,max=11"`
	Balance    float64 `json:"balance" binding:"required,numeric,min=0"`
	IsMerchant bool    `json:"is_merchant" default:"false"`
}

type UserUpdateRequest struct {
	FirstName  string  `json:"first_name" binding:"required,max=100"`
	LastName   string  `json:"last_name" binding:"required,max=100"`
	Balance    float64 `json:"balance" binding:"required,numeric,min=0"`
	IsMerchant bool    `json:"is_merchant" default:"false"`
}
