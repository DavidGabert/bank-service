package account

var _ UseCase = Account{}

type Account struct {
	repository Repository
}

func NewAccountUseCase(repo Repository) *Account {
	return &Account{
		repository: repo,
	}
}
