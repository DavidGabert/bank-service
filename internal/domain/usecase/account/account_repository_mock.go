package account

import (
	"bank-service/internal/domain/entities"
	"context"
	"fmt"
	"github.com/google/uuid"
)

type MockRepository struct {
	CreateFunc          func(ctx context.Context, account *entities.Account) (*entities.Account, error)
	GetAccountFunc      func(ctx context.Context, accountId uuid.UUID) (*entities.Account, error)
	GetAccountByCpfFunc func(ctx context.Context, cpf string) (*entities.Account, error)
	GetListAccountsFunc func(ctx context.Context) ([]*entities.Account, error)
}

func (m *MockRepository) Create(ctx context.Context, account *entities.Account) (*entities.Account, error) {
	if m.CreateFunc != nil {
		return m.CreateFunc(ctx, account)
	}
	return nil, fmt.Errorf("CreateFunc not implemented")
}

func (m *MockRepository) GetAccountById(ctx context.Context, accountId uuid.UUID) (*entities.Account, error) {
	if m.GetAccountFunc != nil {
		return m.GetAccountFunc(ctx, accountId)
	}
	return nil, fmt.Errorf("GetAccountFunc not implemented")
}

func (m *MockRepository) GetAccountByCpf(ctx context.Context, cpf string) (*entities.Account, error) {
	if m.GetAccountByCpfFunc != nil {
		return m.GetAccountByCpfFunc(ctx, cpf)
	}
	return nil, fmt.Errorf("GetAccountByCpfFunc not implemented")
}

func (m *MockRepository) GetListAccounts(ctx context.Context) ([]*entities.Account, error) {
	if m.GetListAccountsFunc != nil {
		return m.GetListAccountsFunc(ctx)
	}
	return nil, fmt.Errorf("GetAccountByCpfFunc not implemented")
}
