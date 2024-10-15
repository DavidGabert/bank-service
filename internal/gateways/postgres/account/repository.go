package account

import (
	"bank-service/internal/domain/usecase/account"
	"database/sql"
)

var _ account.Repository = Repository{}

type Repository struct {
	DB *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		DB: db,
	}
}
