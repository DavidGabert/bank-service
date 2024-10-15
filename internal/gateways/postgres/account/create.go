package account

import (
	"bank-service/internal/domain/entities"
	"context"
)

func (r Repository) Create(ctx context.Context, account *entities.Account) (*entities.Account, error) {
	const query = `
		INSERT INTO account
			(id, name, cpf, secret, balance)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING created_at`

	err := r.DB.QueryRowContext(ctx, query,
		account.Id(),
		account.Name(),
		account.Cpf(),
		account.Secret(),
		account.Balance(),
	).Scan(account.SetCreatedAt)

	if err != nil {
		return nil, err
	}

	return account, nil
}
