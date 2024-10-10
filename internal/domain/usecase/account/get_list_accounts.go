package account

import (
	"bank-service/internal/domain/entities"
	"context"
	"fmt"
)

func (a Account) GetListAccounts(ctx context.Context) ([]*entities.Account, error) {
	accList, err := a.repository.GetListAccounts(ctx)
	if err != nil {
		return nil, fmt.Errorf("get list accounts: %w", err)
	}
	return accList, nil
}
