package transactionManagerMock

import (
	"database/sql"

	transactionManager "github.com/Goboolean/manager-cli/internal/adaptor/transaction-manager"
	"github.com/Goboolean/shared/pkg/resolver"
)

func (t *TransactionManagerMock) Commit() error {

	return nil
}

func (t *TransactionManagerMock) Rollback() error {

	return nil
}

func (t *TransactionManagerMock) TransactionExtractor() transactionManager.TransactionExtractor {
	return t
}

func (t *TransactionManagerMock) TransactionPsql() *sql.Tx {
	return nil
}

func (t *TransactionManagerMock) TransactionMongo() resolver.Transactioner {
	return nil
}
