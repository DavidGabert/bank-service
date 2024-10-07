package usecase

import "bank-service/internal/domain/entities"

func (a Account) Create(input entities.CreateAccountInput) (*entities.Account, error) {
	acc, err := a.repository.Create(entities.NewAccount(input.Name, input.CPF, input.Secret))
	if err != nil {
		return nil, err
	}
	return acc, nil
}
