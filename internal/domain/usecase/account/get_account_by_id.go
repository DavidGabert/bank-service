package account

import (
	"bank-service/internal/domain/entities"
	"context"
	"fmt"
	"github.com/google/uuid"
)

func (a Account) GetAccountById(ctx context.Context, accountId uuid.UUID) (entities.Account, error) {
	acc, err := a.repository.GetAccountById(ctx, accountId)
	if err != nil {
		return entities.Account{}, fmt.Errorf("get account by id: %w", err)
	}
	return acc, nil
}
