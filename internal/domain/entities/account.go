package entities

import (
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

func newId() uuid.UUID {
	return uuid.New()
}

func NewAccount(name string, cpf string, secret string) *Account {
	return &Account{
		id:      newId(),
		name:    name,
		cpf:     cpf,
		secret:  secret,
		balance: 0,
	}
}
