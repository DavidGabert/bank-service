package transfer

import (
	"bank-service/internal/domain/entities"
	"context"
	"fmt"
)

const (
	errMsgTransferAmount = "transfer amount"
)

func (t Transfer) TransferAmount(ctx context.Context, input entities.InputTransfer) error {
	transfer, err := entities.NewTransfer(input.AccountOriginId, input.AccountDestinationId, input.Amount)
	if err != nil {
		return fmt.Errorf("%s: %w", errMsgTransferAmount, err)
	}

	accTransferOrigin, err := t.useCaseAccount.GetAccountById(ctx, input.AccountOriginId)
	if err != nil {
		return fmt.Errorf("%s: %w", errMsgTransferAmount, err)
	}

	accTransferDestin, err := t.useCaseAccount.GetAccountById(ctx, input.AccountDestinationId)
	if err != nil {
		return fmt.Errorf("%s: %w", errMsgTransferAmount, err)
	}

	err = accTransferOrigin.SubtractBalance(input.Amount)
	if err != nil {
		return fmt.Errorf("%s: %w", errMsgTransferAmount, err)
	}

	err = accTransferDestin.AddBalance(input.Amount)
	if err != nil {
		return fmt.Errorf("%s: %w", errMsgTransferAmount, err)
	}

	err = t.repository.PerformTransfer(ctx, &entities.PerformTransferenceInput{
		AccountOrigin:      accTransferOrigin,
		AccountDestination: accTransferDestin,
		Transfer:           transfer,
	})
	if err != nil {
		return fmt.Errorf("%s: %w", errMsgTransferAmount, err)

	}
	return nil
}
