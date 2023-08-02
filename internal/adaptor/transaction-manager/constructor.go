package transactionManager

import (
	"context"
	"database/sql"

	"github.com/Goboolean/shared/pkg/mongo"
	"github.com/Goboolean/shared/pkg/rdbms"
	"github.com/Goboolean/shared/pkg/resolver"
)

type TransactionExtractor interface {
	//Not exposed to the domain layer
	TransactionPsql() *sql.Tx
	TransactionMongo() resolver.Transactioner
}

type TransactionManager struct {
	txPsql  resolver.Transactioner
	txMongo resolver.Transactioner
}

func New(ctx context.Context, psql *rdbms.PSQL, mongo *mongo.DB) (*TransactionManager, error) {
	instance := &TransactionManager{}

	var tx resolver.Transactioner
	var err error

	tx, err = psql.NewTx(ctx)
	if err != nil {
		return nil, err
	}
	instance.txPsql = tx

	tx, err = mongo.NewTx(ctx)
	if err != nil {
		return nil, err
	}
	instance.txMongo = tx

	return instance, nil
}
