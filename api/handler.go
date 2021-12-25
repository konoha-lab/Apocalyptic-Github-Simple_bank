package api

import (
	"database/sql"
	_repo "simple_bank/db/repository"

	_ "github.com/golang/mock/mockgen/model"
)

type Handler interface {
	_repo.Querier
}

type HandlerSQL struct {
	*_repo.Queries
	db *sql.DB
}

func New(db *sql.DB) Handler {
	return &HandlerSQL{
		db:      db,
		Queries: _repo.New(db),
	}
}
