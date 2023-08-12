package metadataRepo

import (
	"errors"

	"github.com/Goboolean/shared/pkg/rdbms"
)

var errExpiredSession = errors.New("session: Commit or Rollback expired session")

type MetadataRepositoryAdaptor struct {
	db *rdbms.PSQL
}

func New(db *rdbms.PSQL) *MetadataRepositoryAdaptor {

	return &MetadataRepositoryAdaptor{
		db: db,
	}
}
