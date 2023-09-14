package transactionCreator

import (
	"github.com/Goboolean/manager-cli/internal/infrastructure/rdbms"
	"github.com/Goboolean/shared/pkg/mongo"
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
