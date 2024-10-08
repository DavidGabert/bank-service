package account

import (
	"bank-service/internal/domain/entities"
	"context"
	"fmt"
)

func (a Account) GetAccountById(ctx context.Context, accountId string) (*entities.Account, error) {
	acc, err := a.repository.GetAccountById(ctx, accountId)
	if err != nil {
		return nil, fmt.Errorf("get account by id: %w", err)
	}
	return acc, nil
}
