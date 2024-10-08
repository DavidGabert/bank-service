package account

import (
	"bank-service/internal/domain/entities"
	"context"
)

type Repository interface {
	Create(ctx context.Context, account *entities.Account) (*entities.Account, error)
}
