package metadataRepo

import (
	"database/sql"
	"errors"
	"os"

	"github.com/Goboolean/shared/pkg/rdbms"
	"github.com/Goboolean/shared/pkg/resolver"
)

var errExpiredSession = errors.New("session: Commit or Rollback expired session")

type MetadataRepositoryAdaptor struct {
	db *rdbms.PSQL

	// Maps session id to query
	queries map[int]*rdbms.Queries

	// Maps session id to transaction
	transactions map[int]*sql.Tx
}

func New() *MetadataRepositoryAdaptor {

	c := &resolver.ConfigMap{
		"HOST":     os.Getenv("PSQL_HOST"),
		"PORT":     os.Getenv("PSQL_PORT"),
		"USER":     os.Getenv("PSQL_USER"),
		"PASSWORD": os.Getenv("PSQL_PASS"),
		"DATABASE": os.Getenv("PSQL_DATABASE"),
	}

	psqlInstance := rdbms.NewDB(c)

	return &MetadataRepositoryAdaptor{
		db:           psqlInstance,
		queries:      make(map[int]*rdbms.Queries),
		transactions: make(map[int]*sql.Tx),
	}
}
