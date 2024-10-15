package ports

import (
	"errors"
)

// account
var (
	ErrCPFAlreadyLinked = errors.New("CPF is already linked to an account")
)

// transfer
var (
	ErrInvalidTransferAmount = errors.New("amount must be greater than zero")
	ErrOrigAccEqualDestAcc   = errors.New("account destination id and origin id cannot be the same")
	ErrInsufficientBalance   = errors.New("account has insufficient balance")
)
