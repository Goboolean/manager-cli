package metadataRepo

import (
	"context"
	"database/sql"
	"os"

	"github.com/Goboolean/shared/pkg/rdbms"
	"github.com/Goboolean/shared/pkg/resolver"
)

type MetadataRepositoryAdaptor struct {
	db  *rdbms.PSQL
	q   *rdbms.Queries
	ctx context.Context
	tx  *sql.Tx
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
		db: psqlInstance,
		q:  psqlInstance.NewQueries(),
	}
}
