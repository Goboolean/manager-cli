package out

import (
	"context"

	transactionManager "github.com/Goboolean/manager-cli/internal/adaptor/transaction-manager"
)

type TransactionCreator interface {
	CreateTransaction(ctx context.Context) TransactorPort
}

type TransactorPort interface {
	Commit() error
	Rollback() error
	TransactionExtractor() transactionManager.TransactionExtractor
}
