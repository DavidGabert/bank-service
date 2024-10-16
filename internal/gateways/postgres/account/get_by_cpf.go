package account

import (
	"bank-service/internal/domain/entities"
	"context"
	"fmt"
	"github.com/google/uuid"
	"time"
)

func (r Repository) GetAccountByCpf(ctx context.Context, cpf string) (entities.Account, error) {

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
			cpf=$1`

	var (
		id                   uuid.UUID
		name, cpfAcc, secret string
		balance              float64
		createdAt            time.Time
	)

	err := r.DB.QueryRowContext(ctx, query, cpf).Scan(
		&id,
		&name,
		&cpfAcc,
		&balance,
		&secret,
		&createdAt,
	)

	if err != nil {
		return entities.Account{}, fmt.Errorf("error getting account by cpf: %w", err)
	}
	return entities.ParseAccount(id, name, cpfAcc, secret, balance, createdAt), nil
}
