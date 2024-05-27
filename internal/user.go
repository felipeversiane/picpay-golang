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
	SetIsMerchant(isMerchant bool)
	GetCreatedAt() time.Time
	GetDocument() string
	SetCreatedAt(createdAt time.Time)
	GetUpdatedAt() time.Time
	SetUpdatedAt(updatedAt time.Time)
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
	isMerchant bool,
	document string,
	balance float64,
) *userDomain {
	return &userDomain{
		id:         uuid.New(),
		email:      email,
		password:   password,
		firstName:  first_name,
		lastName:   last_name,
		isMerchant: isMerchant,
		document:   document,
		balance:    balance,
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

func (u *userDomain) SetIsMerchant(isMerchant bool) {
	u.isMerchant = isMerchant
}

func (u *userDomain) GetCreatedAt() time.Time {
	return u.createdAt
}

func (u *userDomain) SetCreatedAt(createdAt time.Time) {
	u.createdAt = createdAt
}

func (u *userDomain) GetUpdatedAt() time.Time {
	return u.updatedAt
}

func (u *userDomain) SetUpdatedAt(updatedAt time.Time) {
	u.updatedAt = updatedAt
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
