package account

import (
	"bank-service/internal/domain/entities"
	"context"
	"github.com/google/uuid"
	"time"
)

func (r Repository) ListAccounts(ctx context.Context) ([]entities.Account, error) {

	const query = `
		SELECT 
			id,
			name,
			cpf,
			balance,
			created_at
		FROM account
	`

	var (
		id                   uuid.UUID
		name, cpfAcc, secret string
		balance              float64
		createdAt            time.Time
	)

	var accounts []entities.Account

	rows, err := r.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {

		err := rows.Scan(
			&id,
			&name,
			&cpfAcc,
			&balance,
			&createdAt,
		)
		if err != nil {
			return nil, err
		}

		acc := entities.ParseAccount(id, name, cpfAcc, secret, balance, createdAt)
		accounts = append(accounts, acc)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return accounts, nil
}
