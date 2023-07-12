package out

import "github.com/Goboolean/manager-cli/internal/domain/entity"

// TODO: This feature is planned to be implemented later.
type MetadataFetcherPort interface {
	// This method fetches stock overview form remote source
	FetchStockMeta(code string, location string) entity.ProductMeta
	// This method fetches coin overview form remote source
	FetchCoinMeta(code string) entity.ProductMeta
}
