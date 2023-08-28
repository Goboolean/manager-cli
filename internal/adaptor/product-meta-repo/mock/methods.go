package productMetaRepoMock

import (
	"context"

	transactionManager "github.com/Goboolean/manager-cli/internal/adaptor/transaction-manager"
	"github.com/Goboolean/manager-cli/internal/domain/entity"
)

// This method gets unique id of a product which can be hash, UUID and so on...
func (a *MetadataRepositoryAdaptorMock) GetProductId(
	ctx context.Context,
	tx transactionManager.TransactionExtractor,
	code string) (string, error) {

	return "", nil
}

// This method gets full metadata of a product
func (a *MetadataRepositoryAdaptorMock) GetProductMeta(
	ctx context.Context,
	tx transactionManager.TransactionExtractor,
	id string) (entity.ProductMeta, error) {

	return entity.ProductMeta{
		Id:       id,
		Name:     "apple",
		Code:     "AAPL",
		Exchange: "nasdaq",
		Location: "usa",
		Type:     "stock",
	}, nil
}

// This method stores metadata to metadata repository which can be mysql, radius so on...
func (a *MetadataRepositoryAdaptorMock) StoreProductMeta(
	ctx context.Context,
	tx transactionManager.TransactionExtractor,
	meta entity.ProductMeta) error {
	return nil
}

func (a *MetadataRepositoryAdaptorMock) IsProductStored(ctx context.Context, tx transactionManager.TransactionExtractor, id string) (bool, error) {
	if id == "stock.apple.usa" {
		return true, nil
	}
	return false, nil

}

func (a *MetadataRepositoryAdaptorMock) Close() error {
	return nil
}

func (a *MetadataRepositoryAdaptorMock) Ping() error {
	return nil
}

// This method get list of stored product
func (a *MetadataRepositoryAdaptorMock) GetStoredProductList(ctx context.Context, tx transactionManager.TransactionExtractor) ([]string, error) {
	return []string{"stock.apple.usa",
		"stock.nvidia.usa",
		"stock.google.usa",
		"sotck.coca-cola.usa",
		"stock.samsung.kor",
	}, nil
}
