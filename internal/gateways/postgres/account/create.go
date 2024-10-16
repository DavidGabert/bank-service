package account

import (
	"bank-service/internal/domain/entities"
	"context"
	"fmt"
	"time"
)

func (r Repository) Create(ctx context.Context, account entities.Account) (entities.Account, error) {
	const query = `
		INSERT INTO account
			(id, name, cpf, secret, balance)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING created_at`

	var createdAt time.Time

	err := r.DB.QueryRowContext(ctx, query,
		account.Id(),
		account.Name(),
		account.Cpf(),
		account.Secret(),
		account.Balance(),
	).Scan(createdAt)

	account.SetCreatedAt(createdAt)

	if err != nil {
		return entities.Account{}, fmt.Errorf("error creating account: %w", err)
	}

	return account, nil
}
