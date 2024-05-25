package domain

import (
	"crypto/md5"
	"encoding/hex"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID         uuid.UUID `json:"id"`
	Email      string    `json:"email" binding:"required,email"`
	Password   string    `json:"password" binding:"required,min=6,containsany=!@&*%$#"`
	FirstName  string    `json:"first_name" binding:"min=4,max=50"`
	LastName   string    `json:"last_name" binding:"min=4,max=50"`
	Document   string    `json:"balance" binding:"required,numeric,min=0"`
	Balance    float64   `json:"document" binding:"required"`
	IsMerchant bool      `json:"is_merchant" default:"false"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func NewUser(
	email string,
	password string,
	first_name string,
	last_name string,
	isMerchant bool,
	document string,
	balance float64,
) *User {
	return &User{
		ID:         uuid.New(),
		Email:      email,
		Password:   password,
		FirstName:  first_name,
		LastName:   last_name,
		IsMerchant: isMerchant,
		Document:   document,
		Balance:    balance,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
}

func NewUserUpdateDomain(
	first_name string,
	last_name string,
	balance float64,
	isMerchant bool,
) *User {
	return &User{
		FirstName:  first_name,
		LastName:   last_name,
		Balance:    balance,
		IsMerchant: isMerchant,
		UpdatedAt:  time.Now(),
	}
}

func EncryptPassword(password string) string {
	hash := md5.New()
	hash.Write([]byte(password))
	return hex.EncodeToString(hash.Sum(nil))
}
