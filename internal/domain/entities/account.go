package entities

import (
	"bank-service/internal/domain/common/hash"
	"bank-service/internal/domain/ports"
	"github.com/google/uuid"
	"time"
)

type CreateAccountInput struct {
	Name   string
	CPF    string
	Secret string
}

type Account struct {
	id        uuid.UUID
	name      string
	cpf       string
	secret    string
	balance   float64
	createdAt time.Time
}

func ParseAccount(id uuid.UUID, name string, cpf string, secret string, balance float64, createdAt time.Time) Account {
	return Account{
		id:        id,
		name:      name,
		cpf:       cpf,
		secret:    secret,
		balance:   balance,
		createdAt: createdAt,
	}
}

func (a Account) Id() uuid.UUID        { return a.id }
func (a Account) Name() string         { return a.name }
func (a Account) Cpf() string          { return a.cpf }
func (a Account) Secret() string       { return a.secret }
func (a Account) Balance() float64     { return a.balance }
func (a Account) CreatedAt() time.Time { return a.createdAt }

func (a *Account) SetCreatedAt(createdAt time.Time) {
	a.createdAt = createdAt
}

func newId() uuid.UUID {
	return uuid.New()
}

func NewAccount(name string, cpf string, secret string) Account {
	return Account{
		id:      newId(),
		name:    name,
		cpf:     cpf,
		secret:  hash.Hash(secret),
		balance: 0,
	}
}

func (a *Account) AddBalance(amount float64) error {
	if amount <= 0 {
		return ports.ErrInvalidTransferAmount
	}
	a.balance += amount
	return nil
}

func (a *Account) SubtractBalance(amount float64) error {
	if amount <= 0 {
		return ports.ErrInvalidTransferAmount
	} else if amount > a.balance {
		return ports.ErrInsufficientBalance
	}
	a.balance -= amount
	return nil
}
