package account

import (
	"bank-service/internal/domain/entities"
	"fmt"
)

type MockRepository struct {
	CreateFunc func(account *entities.Account) (*entities.Account, error)
}

func (m *MockRepository) Create(account *entities.Account) (*entities.Account, error) {
	if m.CreateFunc != nil {
		return m.CreateFunc(account)
	}
	return nil, fmt.Errorf("CreateFunc not implemented")
}
