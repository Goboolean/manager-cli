package transactionManager

import (
	"database/sql"

	"github.com/Goboolean/shared/pkg/resolver"
)

func (t *TransactionManager) Commit() error {
	var err error

	err = t.txPsql.Commit()
	if err != nil {
		return err
	}

	err = t.txMongo.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (t *TransactionManager) Rollback() error {
	var err error

	err = t.txPsql.Rollback()
	if err != nil {
		return err
	}

	err = t.txMongo.Rollback()
	if err != nil {
		return err
	}

	return nil
}

func (t *TransactionManager) TransactionExtractor() TransactionExtractor {
	return t
}

func (t *TransactionManager) TransactionPsql() *sql.Tx {
	return t.txPsql.Transaction().(*sql.Tx)
}

func (t *TransactionManager) TransactionMongo() resolver.Transactioner {
	return t.txMongo
}
