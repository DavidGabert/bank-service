package transfer

import (
	"bank-service/internal/domain/entities"
	"context"
	"fmt"
)

type MockRepository struct {
	PerformTransferFunc func(ctx context.Context, input entities.PerformTransferenceInput) error
}

func (m *MockRepository) PerformTransfer(ctx context.Context, input entities.PerformTransferenceInput) error {
	if m.PerformTransferFunc != nil {
		return m.PerformTransferFunc(ctx, input)
	}
	return fmt.Errorf("PerformTransferFunc not implemented")
}
