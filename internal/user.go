package domain

import (
	"crypto/md5"
	"encoding/hex"
	"time"

	"github.com/google/uuid"
)

type userDomain struct {
	id         uuid.UUID
	email      string
	password   string
	firstName  string
	lastName   string
	document   string
	balance    float64
	isMerchant bool
	createdAt  time.Time
	updatedAt  time.Time
}

type UserDomainInterface interface {
	GetID() uuid.UUID
	GetEmail() string
	GetIsMerchant() bool
	GetCreatedAt() time.Time
	GetDocument() string
	GetUpdatedAt() time.Time
	GetFirstName() string
	GetLastName() string
	GetPassword() string
	GetBalance() float64
	EncryptPassword()
}

func NewUserDomain(
	email string,
	password string,
	first_name string,
	last_name string,
	document string,
	balance float64,
	isMerchant bool,
) *userDomain {
	return &userDomain{
		id:         uuid.New(),
		email:      email,
		password:   password,
		firstName:  first_name,
		lastName:   last_name,
		document:   document,
		balance:    balance,
		isMerchant: isMerchant,
		createdAt:  time.Now(),
		updatedAt:  time.Now(),
	}
}

func NewUserUpdateDomain(
	first_name string,
	last_name string,
	balance float64,
	isMerchant bool,
) UserDomainInterface {
	return &userDomain{
		firstName:  first_name,
		lastName:   last_name,
		balance:    balance,
		isMerchant: isMerchant,
		updatedAt:  time.Now(),
	}
}

func (u *userDomain) GetID() uuid.UUID {
	return u.id
}

func (u *userDomain) GetEmail() string {
	return u.email
}

func (u *userDomain) GetIsMerchant() bool {
	return u.isMerchant
}

func (u *userDomain) GetCreatedAt() time.Time {
	return u.createdAt
}

func (u *userDomain) GetUpdatedAt() time.Time {
	return u.updatedAt
}

func (u *userDomain) GetFirstName() string {
	return u.firstName
}

func (u *userDomain) GetLastName() string {
	return u.lastName
}

func (u *userDomain) GetPassword() string {
	return u.password
}

func (u *userDomain) GetBalance() float64 {
	return u.balance
}

func (u *userDomain) GetDocument() string {
	return u.document
}

func (u *userDomain) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(u.password))
	u.password = hex.EncodeToString(hash.Sum(nil))
}
