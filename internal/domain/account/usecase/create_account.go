package usecase

import (
	"bank-service/internal/domain/entities"
	"context"
)

func (a Account) Create(ctx context.Context, input entities.CreateAccountInput) (*entities.Account, error) {
	acc, err := a.repository.Create(ctx, entities.NewAccount(input.Name, input.CPF, input.Secret))
	if err != nil {
		return nil, err
	}
	return acc, nil
}
