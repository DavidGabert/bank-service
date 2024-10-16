package entities

import (
	"bank-service/internal/domain/ports"
	"github.com/google/uuid"
	"time"
)

type InputTransfer struct {
	AccountOriginId      uuid.UUID
	AccountDestinationId uuid.UUID
	Amount               float64
}

type PerformTransferenceInput struct {
	AccountOrigin      Account
	AccountDestination Account
	Transfer           Transfer
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

func NewTransfer(accountOriginId uuid.UUID, accountDestinationId uuid.UUID, amount float64) (Transfer, error) {
	newTransfer := Transfer{
		id:                   newId(),
		accountOriginId:      accountOriginId,
		accountDestinationId: accountDestinationId,
		amount:               amount,
	}
	err := newTransfer.validate()
	if err != nil {
		return Transfer{}, err
	}
	return newTransfer, nil
}

func (t Transfer) validate() error {
	if t.amount <= 0 {
		return ports.ErrInvalidTransferAmount
	}
	if t.accountOriginId == t.accountDestinationId {
		return ports.ErrOrigAccEqualDestAcc
	}
	return nil
}
