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
	ID        uuid.UUID
	Name      string
	CPF       string
	Secret    string
	Balance   float64
	CreatedAt time.Time
}

func newId() uuid.UUID {
	return uuid.New()
}

func NewAccount(name string, cpf string, secret string) *Account {
	return &Account{
		ID:      newId(),
		Name:    name,
		CPF:     cpf,
		Secret:  secret,
		Balance: 0,
	}
}
