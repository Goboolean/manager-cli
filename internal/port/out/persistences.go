package out

import (
	"time"

	transactionManager "github.com/Goboolean/manager-cli/internal/adaptor/transaction-manager"
	"github.com/Goboolean/manager-cli/internal/domain/entity"
)

type MetadataRepositoryPort interface {
	// This method get list of stored product
	GetStoredProductList(transaction transactionManager.TransactionExtractor) ([]string, error)
	// This method gets unique id of a product which can be hash, UUID and so on...
	GetProductId(transaction transactionManager.TransactionExtractor, code string) (string, error)
	// This method gets full metadata of a product
	GetProductMeta(transaction transactionManager.TransactionExtractor, id string) (entity.ProductMeta, error)
	// This method stores metadata to metadata repository which can be mysql, radius so on...
	StoreProductMeta(transaction transactionManager.TransactionExtractor, meta entity.ProductMeta) error
}

type StatusPort interface {
	// This method returns status of a product
	GetStatus(id string) (entity.ProductStatus, error)
	// This method changes status of a product by "status" val
	SetStatus(id string, status entity.ProductStatus) error
}

type TradeDumperPort interface {
	// This method dumps trade data of specific product created before time
	DumpProductBefore(id string, outDir string, time time.Time) ([]entity.File, error)
	// This method dumps trade data of specific product created between time
	DumpProductBetween(id string, outDir string, from, to time.Time) ([]entity.File, error)
}
