package transactionCreatorMock

import (
	"context"

	transactionManagerMock "github.com/Goboolean/manager-cli/internal/adaptor/transaction-manager/mock"
	"github.com/Goboolean/manager-cli/internal/port/out"
)

type TransactionFactory struct {
}

func New() *TransactionFactory {
	return &TransactionFactory{}
}

func (f *TransactionFactory) CreateTransaction(ctx context.Context) (out.TransactorPort, error) {
	return transactionManagerMock.New()
}
