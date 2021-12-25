package controller

import (
	"database/sql"
	_repo "simple_bank/db/repository"
)

type Controller struct {
	*_repo.Queries
	db *sql.DB
}

func NewController(db *sql.DB) *Controller {
	return &Controller{
		db:      db,
		Queries: _repo.New(db),
	}
}
