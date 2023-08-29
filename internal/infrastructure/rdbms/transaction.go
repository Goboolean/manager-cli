package rdbms

import (
	"context"
	"database/sql"

	"github.com/Goboolean/shared/pkg/resolver"
)

type Transaction struct {
	tx  *sql.Tx
	ctx context.Context
}

func (t *Transaction) Commit() error {
	return t.tx.Commit()
}

func (t *Transaction) Rollback() error {
	return t.tx.Rollback()
}

func (t *Transaction) Context() context.Context {
	return t.ctx
}

func (t *Transaction) Transaction() interface{} {
	return t.tx
}

func NewTransaction(tx *sql.Tx, ctx context.Context) resolver.Transactioner {
	return &Transaction{tx: tx, ctx: ctx}
}
