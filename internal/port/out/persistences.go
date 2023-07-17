package out

import (
	"time"

	"github.com/Goboolean/manager-cli/internal/domain/entity"
)

type MetadataRepositoryPort interface {
	// This method gets unique id of a product which can be hash, UUID and so on...
	GetProductId(code string) (string, error)
	// This method gets full metadata of a product
	GetProductMeta(id string) (entity.ProductMeta, error)
	// This method stores metadata to metadata repository which can be mysql, radius so on...
	StoreProductMeta(meta entity.ProductMeta) error
}

type StatusPort interface {
	// This method returns status of a product
	GetStatus(id string) (entity.ProductStatus, error)
	// This method changes status of a product by "status" val
	SetStatus(id string, status entity.ProductStatus) error
}

type TradeRepositoryPort interface {
	// This method dumps trade data from trade data repository
	DumpTradeRepo() (entity.FileManager, error)
	// This method dumps trade data created before a specific date
	DumpTradeRepoBefore(date time.Time) (entity.FileManager, error)
	// This method dumps trade data of specific product
	DumpProduct(id string) (entity.FileManager, error)
	// This method dumps trade data of specific product created before time
	DumpProductBefore(id string, time time.Time) (entity.FileManager, error)
}
