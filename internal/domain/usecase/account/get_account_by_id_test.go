package account

import (
	"bank-service/internal/domain/entities"
	"context"
	"errors"
	"github.com/google/uuid"
	"testing"
)

func TestGetAccountById(t *testing.T) {
	t.Parallel()

	errDatabase := errors.New("database error")

	type args struct {
		ctx           context.Context
		accountId     uuid.UUID
		accountEntity entities.Account
	}

	commonArgs := args{
		ctx:           context.Background(),
		accountId:     uuid.New(),
		accountEntity: *entities.NewAccount("John Doe", "843.361.730-36", "SECRET-HASH"),
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
						GetAccountFunc: func(ctx context.Context, id uuid.UUID) (entities.Account, error) {
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
						GetAccountFunc: func(ctx context.Context, id uuid.UUID) (entities.Account, error) {
							return entities.Account{}, errDatabase
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
				t.Errorf("get account by id error = %v, wantErr %v", err, tt.wantError)
			} else if (acc == entities.Account{}) && err == nil {
				t.Errorf("get account by id error, failed to get account")
			}
		})
	}
}
