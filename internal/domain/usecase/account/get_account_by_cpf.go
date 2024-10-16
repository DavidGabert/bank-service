package account

import (
	"bank-service/internal/domain/entities"
	"context"
	"fmt"
)

func (a Account) GetAccountByCpf(ctx context.Context, cpf string) (entities.Account, error) {
	acc, err := a.repository.GetAccountByCpf(ctx, cpf)
	if err != nil {
		return entities.Account{}, fmt.Errorf("get account by cpf: %w", err)
	}
	return acc, nil
}
