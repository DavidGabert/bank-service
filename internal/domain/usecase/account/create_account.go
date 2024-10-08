package account

import (
	"bank-service/internal/domain/entities"
	"context"
	"fmt"
)

func (a Account) Create(ctx context.Context, input entities.CreateAccountInput) (*entities.Account, error) {
	acc, err := a.repository.Create(ctx, entities.NewAccount(input.Name, input.CPF, input.Secret))
	if err != nil {
		return nil, fmt.Errorf("error creating account: %w", err)
	}
	return acc, nil
}
