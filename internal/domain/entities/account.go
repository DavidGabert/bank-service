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

func (a Account) Id() uuid.UUID        { return a.id }
func (a Account) Name() string         { return a.name }
func (a Account) Cpf() string          { return a.cpf }
func (a Account) Secret() string       { return a.secret }
func (a Account) Balance() float64     { return a.balance }
func (a Account) CreatedAt() time.Time { return a.createdAt }

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
