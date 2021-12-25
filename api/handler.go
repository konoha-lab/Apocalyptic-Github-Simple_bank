package api

import (
	"database/sql"
	_repo "simple_bank/db/repository"
)

type Handler struct {
	*_repo.Queries
	db *sql.DB
}

func New(db *sql.DB) *Handler {
	return &Handler{
		db:      db,
		Queries: _repo.New(db),
	}
}
