package account

import (
	"bank-service/internal/domain/entities"
	"context"
	"errors"
	"testing"
)

func TestGetListAccounts(t *testing.T) {
	t.Parallel()

	errDatabase := errors.New("database error")

	type args struct {
		ctx context.Context
	}

	commonArgs := args{
		ctx: context.Background(),
	}

	tests := []struct {
		name      string
		args      args
		setup     func(t *testing.T) Account
		want      []*entities.Account
		wantError error
	}{
		{
			name: "should return list of accounts without error",
			args: commonArgs,
			setup: func(t *testing.T) Account {
				return Account{
					repository: &MockRepository{
						ListAccountsFunc: func(ctx context.Context) ([]entities.Account, error) {
							return []entities.Account{
								entities.NewAccount("John Doe", "123.456.789-00", "SECRET-HASH-1"),
								entities.NewAccount("Jane Doe", "987.654.321-00", "SECRET-HASH-2"),
							}, nil
						},
					},
				}
			},
			want:      []*entities.Account{},
			wantError: nil,
		},
		{
			name: "should return database error when trying to get list of accounts",
			args: commonArgs,
			setup: func(t *testing.T) Account {
				return Account{
					repository: &MockRepository{
						ListAccountsFunc: func(ctx context.Context) ([]entities.Account, error) {
							return nil, errDatabase
						},
					},
				}
			},
			want:      nil,
			wantError: errDatabase,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			accounts, err := tt.setup(t).ListAccounts(tt.args.ctx)

			if err != nil && !errors.Is(err, tt.wantError) {
				t.Errorf("get list account error = %v, wantErr %v", err, tt.wantError)
			}
			if err == nil && (accounts == nil || len(accounts) == 0) {
				t.Errorf("get list account error returned nil or empty list of accounts")
				return
			}
		})
	}
}
