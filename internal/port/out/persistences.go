package out

import (
	"time"

	"github.com/Goboolean/manager-cli/internal/domain/entity"
)

type MetadataRepositoryPort interface {
	TransactorPort

	// This method gets unique id of a product which can be hash, UUID and so on...
	GetProductId(session entity.Session, code string) (string, error)
	// This method gets full metadata of a product
	GetProductMeta(session entity.Session, id string) (entity.ProductMeta, error)
	// This method stores metadata to metadata repository which can be mysql, radius so on...
	StoreProductMeta(session entity.Session, meta entity.ProductMeta) error
}

type StatusPort interface {
	// This method returns status of a product
	GetStatus(id string) (entity.ProductStatus, error)
	// This method changes status of a product by "status" val
	SetStatus(id string, status entity.ProductStatus) error
}

type TradeRepositoryPort interface {
	TransactorPort

	// This method dumps trade data from trade data repository
	DumpTradeRepo(session entity.Session) (entity.FileManager, error)
	// This method dumps trade data created before a specific date
	DumpTradeRepoBefore(session entity.Session, date time.Time) (entity.FileManager, error)
	// This method dumps trade data of specific product
	DumpProduct(session entity.Session, id string) (entity.FileManager, error)
	// This method dumps trade data of specific product created before time
	DumpProductBefore(session entity.Session, id string, time time.Time) (entity.FileManager, error)
}
