package transfer

import "bank-service/internal/domain/usecase/account"

type Transfer struct {
	repository     Repository
	useCaseAccount account.Account
}

func NewTransferUseCase(repo Repository, useCaseAccount account.Account) Transfer {
	return Transfer{
		repository:     repo,
		useCaseAccount: useCaseAccount,
	}
}
