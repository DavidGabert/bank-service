package transfer

import "bank-service/internal/domain/usecase/account"

var _ UseCase = Transfer{}

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
