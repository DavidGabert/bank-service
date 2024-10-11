package account

import (
	"bank-service/internal/domain/entities"
	"context"
	"fmt"
)

func (a Account) ListAccounts(ctx context.Context) ([]entities.Account, error) {
	accList, err := a.repository.ListAccounts(ctx)
	if err != nil {
		return nil, fmt.Errorf("list accounts: %w", err)
	}
	return accList, nil
}
