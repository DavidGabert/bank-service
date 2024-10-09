package account

import (
	"bank-service/internal/domain/entities"
	"context"
	"errors"
	"testing"
)

func TestGetAccountById(t *testing.T) {
	t.Parallel()

	errDatabase := errors.New("database error")

	type args struct {
		ctx           context.Context
		accountId     string
		accountEntity *entities.Account
	}

	commonArgs := args{
		ctx:           context.Background(),
		accountId:     "a8f5f167f44f4964e6c998dee827110c",
		accountEntity: entities.NewAccount("John Doe", "843.361.730-36", "SECRET-HASH"),
	}

	tests := []struct {
		name      string
		args      args
		setup     func(t *testing.T) Account
		wantError error
	}{
		{
			name: "should return nil error and find an account",
			args: commonArgs,

			setup: func(t *testing.T) Account {
				return Account{
					repository: &MockRepository{
						GetAccountFunc: func(ctx context.Context, id string) (*entities.Account, error) {
							return commonArgs.accountEntity, nil
						},
					},
				}
			},
			wantError: nil,
		},
		{
			name: "should return an database error when trying to find an account",
			args: commonArgs,
			setup: func(t *testing.T) Account {
				return Account{
					repository: &MockRepository{
						GetAccountFunc: func(ctx context.Context, id string) (*entities.Account, error) {
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

			acc, err := tt.setup(t).GetAccountById(tt.args.ctx, tt.args.accountId)
			if err != nil && !errors.Is(err, tt.wantError) {
				t.Errorf("GetAccountById() error = %v, wantErr %v", err, tt.wantError)
			} else if acc == nil && err == nil {
				t.Errorf("GetAccountById() error, failed to get account")
			}
		})
	}
}
