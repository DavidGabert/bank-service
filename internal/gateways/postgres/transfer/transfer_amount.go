package transfer

import (
	"bank-service/internal/domain/entities"
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"
)

// TODO: REFACTOR!!! SEGREGATE METHODS AND RESPONSIBILITIES. BY NOW EVERYTHING IS TOGETHER TO MAKE EASY DO AN TX ROLLBACK

const (
	errMsgTransferAmount = "transfer amount error"
	updateBalanceQuery   = `
		UPDATE account
		SET balance = $1
		WHERE id = $2`
	saveTransferQuery = `
		INSERT INTO transfer
			(id, account_origin_id, account_destination_id, amount)
		VALUES ($1, $2, $3, $4)
		RETURNING created_at`
)

func (r Repository) PerformTransfer(ctx context.Context, input entities.PerformTransferenceInput) error {
	tx, err := r.DB.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}
	defer r.handleTransactionRollback(tx, &err)

	if err = r.updateAccountBalance(ctx, tx, input.AccountOrigin); err != nil {
		return r.wrapTransferMessageError(err)
	}

	if err = r.updateAccountBalance(ctx, tx, input.AccountDestination); err != nil {
		return r.wrapTransferMessageError(err)
	}

	if err = r.saveTransfer(ctx, tx, input.Transfer); err != nil {
		return r.wrapTransferMessageError(err)
	}

	if err = tx.Commit(); err != nil {
		return r.wrapTransferMessageError(err)
	}

	return nil
}

func (r Repository) handleTransactionRollback(tx *sql.Tx, err *error) {
	if *err != nil {
		if rollErr := tx.Rollback(); rollErr != nil {
			log.Println("rollback error:", rollErr)
		}
	}
}

func (r Repository) wrapTransferMessageError(err error) error {
	return fmt.Errorf("%s: %w", errMsgTransferAmount, err)
}

func (r Repository) updateAccountBalance(ctx context.Context, tx *sql.Tx, acc entities.Account) error {
	_, err := tx.ExecContext(ctx, updateBalanceQuery, acc.Balance(), acc.Id())
	return err
}

func (r Repository) saveTransfer(ctx context.Context, tx *sql.Tx, trans entities.Transfer) error {
	var createdAt time.Time

	err := tx.QueryRowContext(
		ctx,
		saveTransferQuery,
		trans.Id(),
		trans.AccountOriginId(),
		trans.AccountDestinationId(),
		trans.Amount(),
	).Scan(&createdAt)

	if err != nil {
		return err
	}

	trans.SetCreatedAt(createdAt)
	return nil
}
