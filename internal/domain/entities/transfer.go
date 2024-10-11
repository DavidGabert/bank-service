package entities

import (
	"errors"
	"github.com/google/uuid"
	"time"
)

type InputTransfer struct {
	AccountOriginId      uuid.UUID
	AccountDestinationId uuid.UUID
	Amount               float64
}

type PerformTransferenceInput struct {
	AccountOrigin      *Account
	AccountDestination *Account
	Transfer           *Transfer
}

type Transfer struct {
	id                   uuid.UUID
	accountOriginId      uuid.UUID
	accountDestinationId uuid.UUID
	amount               float64
	createdAt            time.Time
}

func (t Transfer) AccountOriginId() uuid.UUID      { return t.accountOriginId }
func (t Transfer) AccountDestinationId() uuid.UUID { return t.accountDestinationId }

func NewTransfer(accountOriginId uuid.UUID, accountDestinationId uuid.UUID, amount float64) (*Transfer, error) {
	newTransfer := &Transfer{
		id:                   newId(),
		accountOriginId:      accountOriginId,
		accountDestinationId: accountDestinationId,
		amount:               amount,
	}
	err := newTransfer.validate()
	if err != nil {
		return nil, err
	}
	return newTransfer, nil
}

func (t Transfer) validate() error {
	if t.amount <= 0 {
		return errors.New("amount must be greater than zero")
	}
	if t.accountOriginId == t.accountDestinationId {
		return errors.New("account destination id and origin id cannot be the same")
	}
	return nil
}
