package metadataRepo

import (
	"database/sql"
	"errors"

	"github.com/Goboolean/shared/pkg/rdbms"
)

var errExpiredSession = errors.New("session: Commit or Rollback expired session")

type MetadataRepositoryAdaptor struct {
	db *rdbms.PSQL

	// Maps session id to query
	queries map[int]*rdbms.Queries

	// Maps session id to transaction
	transactions map[int]*sql.Tx
}

func New(db *rdbms.PSQL) *MetadataRepositoryAdaptor {

	return &MetadataRepositoryAdaptor{
		db:           db,
		queries:      make(map[int]*rdbms.Queries),
		transactions: make(map[int]*sql.Tx),
	}
}
