package account

import (
	"bank-service/internal/domain/common/hash"
	"bank-service/internal/domain/entities"
	"context"
	"fmt"
)

var (
	ErrCPFAlreadyLinked = fmt.Errorf("CPF is already linked to an account")
)

func (a Account) Create(ctx context.Context, input entities.CreateAccountInput) (*entities.Account, error) {
	acc, err := a.repository.GetAccountByCpf(ctx, input.CPF)
	if err != nil {
		return nil, fmt.Errorf("get account by cpf: %w", err)
	}
	if acc != nil {
		return nil, ErrCPFAlreadyLinked
	}
	hashSecret := hash.Hash(input.Secret)
	acc, err = a.repository.Create(ctx, entities.NewAccount(input.Name, input.CPF, hashSecret))
	if err != nil {
		return nil, fmt.Errorf("error creating account: %w", err)
	}
	return acc, nil
}
