package ctrl

import (
	"context"
	"database/sql"
	"fmt"
	_repo "simple_bank/db/repository"
)

type HandlerSQL struct {
	*_repo.Queries
	Db *sql.DB
}

func (handler *HandlerSQL) execTx(ctx context.Context, fn func(*_repo.Queries) error) error {
	tx, err := handler.Db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := _repo.New(tx)

	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}
