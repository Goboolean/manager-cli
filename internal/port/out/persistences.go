package out

import (
	"time"

	"github.com/Goboolean/manager-cli/internal/domain/entity"
)

type MetadataRepositoryPort interface {
	// This method gets unique identifier of a product which can be hash, UUID and so on...
	GetProductIdentifier(code string) (string, error)
	// This method gets full metadata of a product
	GetProductMeta(identifier string) (entity.ProductMeta, error)
	// This method stores metadata to metadata repository which can be mysql, radius so on...
	StoreProductMeta(meta entity.ProductMeta) error
}

type StatusPort interface {
	// This method returns status of a product
	GetStatus(identifier string) entity.ProductStatus
	// This method changes status of a product by "status" val
	ChangeStatus(productIdentifier string, status entity.ProductStatus) error
}

type TradeRepositoryPort interface {
	// This method dumps trade data from trade data repository
	DumpTradeRepo() entity.FileManager
	// This method dumps trade data created before a specific date
	DumpTradeRepoBefore(date time.Time) entity.FileManager
	// This method dumps trade data of specific product
	DumpProduct(identifier string) entity.FileManager
}
