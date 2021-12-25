package api

import (
	"context"
	"database/sql"
	_ctrl "simple_bank/api/controller"
	_repo "simple_bank/db/repository"

	_ "github.com/golang/mock/mockgen/model"
)

type Handler interface {
	_repo.Querier
	ST_TransferTx(ctx context.Context, arg _ctrl.TransferTxParams) (_ctrl.TransferTxResult, error)
}

func New(db *sql.DB) Handler {
	return &_ctrl.HandlerSQL{
		Db:      db,
		Queries: _repo.New(db),
	}
}
