package metadataRepo

import (
	"github.com/Goboolean/manager-cli/infrastructure/rdbms"
)

type MetadataRepositoryAdaptor struct {
	db *rdbms.PSQL
}

func New(db *rdbms.PSQL) *MetadataRepositoryAdaptor {
	return &MetadataRepositoryAdaptor{
		db: db,
	}
}
