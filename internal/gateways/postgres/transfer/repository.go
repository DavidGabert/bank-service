package transfer

import (
	"bank-service/internal/domain/usecase/transfer"
	"database/sql"
)

var _ transfer.Repository = Repository{}

type Repository struct {
	DB *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		DB: db,
	}
}
