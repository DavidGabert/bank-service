package account

import (
	"bank-service/internal/domain/entities"
	"context"
	"fmt"
	"github.com/google/uuid"
)

type MockRepository struct {
	CreateFunc          func(ctx context.Context, account entities.Account) (entities.Account, error)
	GetAccountFunc      func(ctx context.Context, accountId uuid.UUID) (entities.Account, error)
	GetAccountByCpfFunc func(ctx context.Context, cpf string) (entities.Account, error)
	ListAccountsFunc    func(ctx context.Context) ([]entities.Account, error)
}

func (m *MockRepository) Create(ctx context.Context, account entities.Account) (entities.Account, error) {
	if m.CreateFunc != nil {
		return m.CreateFunc(ctx, account)
	}
	return entities.Account{}, fmt.Errorf("CreateFunc not implemented")
}

func (m *MockRepository) GetAccountById(ctx context.Context, accountId uuid.UUID) (entities.Account, error) {
	if m.GetAccountFunc != nil {
		return m.GetAccountFunc(ctx, accountId)
	}
	return entities.Account{}, fmt.Errorf("GetAccountFunc not implemented")
}

func (m *MockRepository) GetAccountByCpf(ctx context.Context, cpf string) (entities.Account, error) {
	if m.GetAccountByCpfFunc != nil {
		return m.GetAccountByCpfFunc(ctx, cpf)
	}
	return entities.Account{}, fmt.Errorf("GetAccountByCpfFunc not implemented")
}

func (m *MockRepository) ListAccounts(ctx context.Context) ([]entities.Account, error) {
	if m.ListAccountsFunc != nil {
		return m.ListAccountsFunc(ctx)
	}
	return nil, fmt.Errorf("GetAccountByCpfFunc not implemented")
}
