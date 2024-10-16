package account

import (
	"bank-service/internal/domain/entities"
	"context"
	"github.com/google/uuid"
)

type Repository interface {
	Create(ctx context.Context, account *entities.Account) (*entities.Account, error)
	GetAccountById(ctx context.Context, accountId uuid.UUID) (entities.Account, error)
	GetAccountByCpf(ctx context.Context, cpf string) (entities.Account, error)
	ListAccounts(ctx context.Context) ([]entities.Account, error)
}
