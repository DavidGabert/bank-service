package account

import (
	"bank-service/internal/domain/entities"
	"context"
)

type Repository interface {
	Create(ctx context.Context, account *entities.Account) (*entities.Account, error)
	GetAccountById(ctx context.Context, accountId string) (*entities.Account, error)
	GetAccountByCpf(ctx context.Context, cpf string) (*entities.Account, error)
	ListAccounts(ctx context.Context) ([]entities.Account, error)
}
