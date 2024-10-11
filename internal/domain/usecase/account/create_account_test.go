package account

import (
	"bank-service/internal/domain/entities"
	"bank-service/internal/domain/ports"
	"context"
	"errors"
	"testing"
)

func TestCreateAccount(t *testing.T) {
	t.Parallel()

	errDatabase := errors.New("database error")

	type args struct {
		ctx            context.Context
		input          entities.CreateAccountInput
		existingCpfAcc *entities.Account
	}

	commonArgs := args{
		ctx: context.Background(),
		input: entities.CreateAccountInput{
			Name:   "Jon Doe",
			CPF:    "843.361.730-36",
			Secret: "SECRET-HASH",
		},
		existingCpfAcc: entities.NewAccount("Jane Doe", "843.361.730-36", "SECRET-HASH-JANE"),
	}

	tests := []struct {
		name      string
		args      args
		setup     func(t *testing.T) Account
		wantError error
	}{
		{
			name: "should return nil error and create an account",
			args: commonArgs,
			setup: func(t *testing.T) Account {
				return Account{
					repository: &MockRepository{
						CreateFunc: func(ctx context.Context, account *entities.Account) (*entities.Account, error) {
							return entities.NewAccount(account.Name(), account.Cpf(), account.Secret()), nil
						},
						GetAccountByCpfFunc: func(ctx context.Context, cpf string) (*entities.Account, error) {
							return nil, nil
						},
					},
				}
			},
			wantError: nil,
		},
		{
			name: "should return an database error when trying to create an account",
			args: commonArgs,
			setup: func(t *testing.T) Account {
				return Account{
					repository: &MockRepository{
						CreateFunc: func(ctx context.Context, account *entities.Account) (*entities.Account, error) {
							return nil, errDatabase
						},
						GetAccountByCpfFunc: func(ctx context.Context, cpf string) (*entities.Account, error) {
							return nil, nil
						},
					},
				}
			},
			wantError: errDatabase,
		},
		{
			name: "should return error cpf already linked to one account",
			args: commonArgs,
			setup: func(t *testing.T) Account {
				return Account{
					repository: &MockRepository{
						CreateFunc: func(ctx context.Context, account *entities.Account) (*entities.Account, error) {
							return entities.NewAccount(account.Name(), account.Cpf(), account.Secret()), nil
						},
						GetAccountByCpfFunc: func(ctx context.Context, cpf string) (*entities.Account, error) {
							return commonArgs.existingCpfAcc, nil
						},
					},
				}
			},
			wantError: ports.ErrCPFAlreadyLinked,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			acc, err := tt.setup(t).Create(tt.args.ctx, tt.args.input)
			if err != nil && !errors.Is(err, tt.wantError) {
				t.Errorf("create account failed test = %v, wantErr %v", err, tt.wantError)
			} else if acc == nil && err == nil {
				t.Errorf("create account failed test, account not created")
			}
		})
	}
}
