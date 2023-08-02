package transactionFactory

import (
	"context"

	transactionManager "github.com/Goboolean/manager-cli/internal/adaptor/transaction-manager"
	"github.com/Goboolean/manager-cli/internal/port/out"
	"github.com/Goboolean/shared/pkg/mongo"
	"github.com/Goboolean/shared/pkg/rdbms"
)

type TransactionFactory struct {
	psqldb  *rdbms.PSQL
	mongodb *mongo.DB
}

func New(psqlDb *rdbms.PSQL, mongoDb *mongo.DB) *TransactionFactory {
	return &TransactionFactory{
		psqldb:  psqlDb,
		mongodb: mongoDb,
	}
}

func (f *TransactionFactory) BuildTransaction(ctx context.Context) (out.Transactor, error) {
	return transactionManager.New(ctx, f.psqldb, f.mongodb)
}
