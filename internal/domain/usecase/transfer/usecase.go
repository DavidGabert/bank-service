package transfer

import (
	"bank-service/internal/domain/entities"
	"context"
)

type UseCase interface {
	TransferAmount(ctx context.Context, transfer entities.Transfer) error
}
