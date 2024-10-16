package account

import (
	"bank-service/internal/domain/entities"
	"bank-service/internal/domain/ports"
	"context"
	"fmt"
)

func (a Account) Create(ctx context.Context, input entities.CreateAccountInput) (entities.Account, error) {
	acc, err := a.repository.GetAccountByCpf(ctx, input.CPF)
	if err != nil {
		return entities.Account{}, fmt.Errorf("get account by cpf: %w", err)
	}
	if (acc != entities.Account{}) {
		return entities.Account{}, ports.ErrCPFAlreadyLinked
	}
	acc, err = a.repository.Create(ctx, entities.NewAccount(input.Name, input.CPF, input.Secret))
	if err != nil {
		return entities.Account{}, fmt.Errorf("error creating account: %w", err)
	}
	return acc, nil
}
