package usecase

import (
	"bank-service/internal/domain/entities"
)

type UseCase interface {
	Create(input entities.CreateAccountInput) (*entities.Account, error)
}
