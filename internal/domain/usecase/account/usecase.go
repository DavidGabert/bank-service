package account

import (
	"bank-service/internal/domain/entities"
	"context"
)

type UseCase interface {
	Create(ctx context.Context, input entities.CreateAccountInput) (*entities.Account, error)
	GetAccountById(ctx context.Context, accountId string) (*entities.Account, error)
	GetAccountByCpf(ctx context.Context, cpf string) (*entities.Account, error)
}
