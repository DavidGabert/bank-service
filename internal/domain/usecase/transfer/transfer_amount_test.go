package transfer

import (
	"bank-service/internal/domain/entities"
	"bank-service/internal/domain/ports"
	"bank-service/internal/domain/usecase/account"
	"context"
	"errors"
	"github.com/google/uuid"
	"testing"
)

func TestTransfer(t *testing.T) {

	originUuidAcc := uuid.New()
	destinUuidAcc := uuid.New()
	originAccBalance := 1200.00
	errDatabase := errors.New("database error")

	type args struct {
		ctx   context.Context
		input entities.InputTransfer
	}

	commonArgs := args{
		ctx: context.Background(),
		input: entities.InputTransfer{
			AccountOriginId:      originUuidAcc,
			AccountDestinationId: destinUuidAcc,
			Amount:               1000.0,
		},
	}

	tests := []struct {
		name      string
		args      args
		setup     func(t *testing.T) Transfer
		want      []*entities.Account
		wantError error
	}{
		{
			name: "should return nil error and transfer with success",
			args: commonArgs,
			setup: func(t *testing.T) Transfer {
				return NewTransferUseCase(&MockRepository{
					PerformTransferFunc: func(ctx context.Context, input entities.PerformTransferenceInput) error {
						return nil
					}}, *account.NewAccountUseCase(&account.MockRepository{
					GetAccountFunc: func(ctx context.Context, accountId uuid.UUID) (entities.Account, error) {
						if accountId == originUuidAcc {
							acc := entities.NewAccount("Account Test Origin", "106.435.680-00", "SECRET-HASH-ORI")
							err := acc.AddBalance(originAccBalance)
							if err != nil {
								return entities.Account{}, err
							}
							return acc, nil
						} else if accountId == destinUuidAcc {
							return entities.NewAccount("Account Test Destination", "073.294.310-87", "SECRET-HASH-DES"), nil
						}
						return entities.Account{}, nil
					},
				}))
			},
			wantError: nil,
		},
		{
			name: "should return ErrInsufficientBalance error",
			args: commonArgs,
			setup: func(t *testing.T) Transfer {
				return NewTransferUseCase(&MockRepository{
					PerformTransferFunc: func(ctx context.Context, input entities.PerformTransferenceInput) error {
						return nil
					}}, *account.NewAccountUseCase(&account.MockRepository{
					GetAccountFunc: func(ctx context.Context, accountId uuid.UUID) (entities.Account, error) {
						if accountId == originUuidAcc {
							return entities.NewAccount("Account Test Origin", "106.435.680-00", "SECRET-HASH-ORI"), nil
						} else if accountId == destinUuidAcc {
							return entities.NewAccount("Account Test Destination", "073.294.310-87", "SECRET-HASH-DES"), nil
						}
						return entities.Account{}, nil
					},
				}))
			},
			wantError: ports.ErrInsufficientBalance,
		},
		{
			name: "should return an databaseErr error",
			args: commonArgs,
			setup: func(t *testing.T) Transfer {
				return NewTransferUseCase(&MockRepository{
					PerformTransferFunc: func(ctx context.Context, input entities.PerformTransferenceInput) error {
						return nil
					}}, *account.NewAccountUseCase(&account.MockRepository{
					GetAccountFunc: func(ctx context.Context, accountId uuid.UUID) (entities.Account, error) {
						if accountId == originUuidAcc {
							return entities.NewAccount("Account Test Origin", "106.435.680-00", "SECRET-HASH-ORI"), nil
						} else if accountId == destinUuidAcc {
							return entities.Account{}, errDatabase
						}
						return entities.Account{}, nil
					},
				}))
			},
			wantError: errDatabase,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			err := tt.setup(t).TransferAmount(tt.args.ctx, tt.args.input)
			if err != nil && !errors.Is(err, tt.wantError) {
				t.Errorf("transfer amount error = %v, wantErr %v", err, tt.wantError)
			}
		})
	}
}
