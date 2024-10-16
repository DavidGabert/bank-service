package transfer

import (
	"bank-service/internal/domain/entities"
	"context"
)

type Repository interface {
	PerformTransfer(ctx context.Context, input entities.PerformTransferenceInput) error
}
