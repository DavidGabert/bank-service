package usecase

import (
	account2 "bank-service/internal/domain/account"
	"bank-service/internal/domain/entities"
	"fmt"
	"github.com/google/uuid"
	"testing"
)

func TestCreate(t *testing.T) {
	t.Parallel()
	mockRepo := account2.MockRepository{}
	account := Account{repository: &mockRepo}
	input := entities.CreateAccountInput{
		Name:   "John Doe",
		CPF:    "12345678900",
		Secret: "secret",
	}

	mockRepo.CreateFunc = func(acc *entities.Account) (*entities.Account, error) {
		return &entities.Account{
			ID:     uuid.New(),
			Name:   acc.Name,
			CPF:    acc.CPF,
			Secret: acc.Secret,
		}, nil
	}

	result, err := account.Create(input)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if result == nil {
		t.Error("Expected a non-nil result")
	}

	if result.Name != input.Name {
		t.Errorf("Expected name %s, got %s", input.Name, result.Name)
	}

	if result.CPF != input.CPF {
		t.Errorf("Expected CPF %s, got %s", input.CPF, result.CPF)
	}

	if result.Secret != input.Secret {
		t.Errorf("Expected secret %s, got %s", input.Secret, result.Secret)
	}
}

func TestCreate_Error(t *testing.T) {
	t.Parallel()
	mockRepo := &account2.MockRepository{}
	account := Account{repository: mockRepo}
	input := entities.CreateAccountInput{
		Name:   "John Doe",
		CPF:    "12345678900",
		Secret: "secret123",
	}

	mockRepo.CreateFunc = func(acc *entities.Account) (*entities.Account, error) {
		return nil, fmt.Errorf("database error")
	}

	result, err := account.Create(input)

	if err == nil {
		t.Error("Expected an error, got nil")
	}

	if result != nil {
		t.Error("Expected a nil result")
	}
}
