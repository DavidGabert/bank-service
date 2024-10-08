package account

import (
	"bank-service/internal/domain/entities"
	"context"
	"fmt"
)

type MockRepository struct {
	CreateFunc     func(ctx context.Context, account *entities.Account) (*entities.Account, error)
	GetAccountFunc func(ctx context.Context, id string) (*entities.Account, error)
}

func (m *MockRepository) Create(ctx context.Context, account *entities.Account) (*entities.Account, error) {
	if m.CreateFunc != nil {
		return m.CreateFunc(ctx, account)
	}
	return nil, fmt.Errorf("CreateFunc not implemented")
}

func (m *MockRepository) GetAccountById(ctx context.Context, id string) (*entities.Account, error) {
	if m.GetAccountFunc != nil {
		return m.GetAccountFunc(ctx, id)
	}
	return nil, fmt.Errorf("GetAccountFunc not implemented")
}
