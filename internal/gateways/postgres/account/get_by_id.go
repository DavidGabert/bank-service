package account

import (
	"bank-service/internal/domain/entities"
	"context"
	"fmt"
	"github.com/google/uuid"
	"time"
)

func (r Repository) GetAccountById(ctx context.Context, accountId uuid.UUID) (entities.Account, error) {
	const query = `
		SELECT
			id,
			name,
			cpf,
			balance,
			secret,
			created_at
		FROM
			account
		WHERE
			id=$1`

	var (
		id                   uuid.UUID
		name, cpfAcc, secret string
		balance              float64
		createdAt            time.Time
	)

	err := r.DB.QueryRowContext(ctx, query, accountId).Scan(
		&id,
		&name,
		&cpfAcc,
		&balance,
		&secret,
		&createdAt,
	)

	if err != nil {
		return entities.Account{}, fmt.Errorf("error getting account by id: %w", err)
	}
	return entities.ParseAccount(id, name, cpfAcc, secret, balance, createdAt), nil
}
