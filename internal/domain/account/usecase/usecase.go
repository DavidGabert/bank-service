package usecase

import (
	"bank-service/internal/domain/entities"
	"context"
)

type UseCase interface {
	Create(ctx context.Context, input entities.CreateAccountInput) (*entities.Account, error)
}
