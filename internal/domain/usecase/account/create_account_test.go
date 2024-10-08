package account

import (
	"bank-service/internal/domain/entities"
	"context"
	"errors"
	"testing"
)

func TestCreateAccount(t *testing.T) {
	t.Parallel()

	errDatabase := errors.New("database error")

	type args struct {
		ctx   context.Context
		input entities.CreateAccountInput
	}

	commonArgs := args{
		ctx: context.Background(),
		input: entities.CreateAccountInput{
			Name:   "Jon Doe",
			CPF:    "843.361.730-36",
			Secret: "SECRET-HASH",
		},
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
					},
				}
			},
			wantError: errDatabase,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			acc, err := tt.setup(t).Create(tt.args.ctx, tt.args.input)
			if err != nil && !errors.Is(err, tt.wantError) {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantError)
			} else if acc == nil && err == nil {
				t.Errorf("Create() error, Account not created")
			}
		})
	}
}
