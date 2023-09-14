package transactionCreator

import (
	"context"

	transactionManager "github.com/Goboolean/manager-cli/internal/adaptor/transaction-manager"
	"github.com/Goboolean/manager-cli/internal/port/out"
)

func (f *TransactionFactory) CreateTransaction(ctx context.Context) (out.TransactorPort, error) {
	return transactionManager.New(ctx, f.psqldb, f.mongodb)
}
