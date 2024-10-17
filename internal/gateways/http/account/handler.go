package account

import (
	"bank-service/internal/domain/usecase/account"
)

type Handler struct {
	AccountUc account.UseCase
}
