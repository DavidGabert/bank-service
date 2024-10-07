package account

import "bank-service/internal/domain/entities"

type Repository interface {
	Create(account *entities.Account) (*entities.Account, error)
}
