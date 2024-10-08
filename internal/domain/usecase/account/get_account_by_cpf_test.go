package account

import (
	"bank-service/internal/domain/entities"
	"context"
	"errors"
	"testing"
)

func TestGetAccountByCpf(t *testing.T) {
	t.Parallel()

	errDatabase := errors.New("database error")

	type args struct {
		ctx           context.Context
		cpf           string
		accountEntity *entities.Account
	}

	commonArgs := args{
		ctx:           context.Background(),
		cpf:           "843.361.730-36",
		accountEntity: entities.NewAccount("John Doe", "843.361.730-36", "SECRET-HASH"),
	}

	tests := []struct {
		name      string
		args      args
		setup     func(t *testing.T) Account
		wantError error
	}{
		{
			name: "should return nil error and find an account by cpf",
			args: commonArgs,

			setup: func(t *testing.T) Account {
				return Account{
					repository: &MockRepository{
						GetAccountByCpfFunc: func(ctx context.Context, cpf string) (*entities.Account, error) {
							return commonArgs.accountEntity, nil
						},
					},
				}
			},
			wantError: nil,
		},
		{
			name: "should return an database error when trying to find an account by cpf",
			args: commonArgs,
			setup: func(t *testing.T) Account {
				return Account{
					repository: &MockRepository{
						GetAccountByCpfFunc: func(ctx context.Context, cpf string) (*entities.Account, error) {
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

			acc, err := tt.setup(t).GetAccountByCpf(tt.args.ctx, tt.args.cpf)
			if err != nil && !errors.Is(err, tt.wantError) {
				t.Errorf("GetAccountByCpf() error = %v, wantErr %v", err, tt.wantError)
			} else if acc == nil && err == nil {
				t.Errorf("GetAccountByCpf() error, failed to get account")
			}
		})
	}
}
