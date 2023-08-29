package productMetaRepo

import (
	"github.com/Goboolean/manager-cli/internal/infrastructure/rdbms"
)

type MetadataRepositoryAdaptor struct {
	db *rdbms.PSQL
}

func New(db *rdbms.PSQL) *MetadataRepositoryAdaptor {
	return &MetadataRepositoryAdaptor{
		db: db,
	}
}
