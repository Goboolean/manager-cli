package metadataRepo

import (
	"context"
	"database/sql"

	"github.com/Goboolean/manager-cli/infrastructure/rdbms"
	transactionManager "github.com/Goboolean/manager-cli/internal/adaptor/transaction-manager"
	"github.com/Goboolean/manager-cli/internal/domain/entity"
)

// This method gets unique id of a product which can be hash, UUID and so on...
func (a *MetadataRepositoryAdaptor) GetProductId(
	ctx context.Context,
	tx transactionManager.TransactionExtractor,
	code string) (string, error) {

	q := rdbms.NewQueries(a.db).WithTx(tx.TransactionPsql())
	result, err := q.GetProductIdBySymbol(ctx, code)

	if err != nil {
		return "", err
	}

	return result, err

}

// This method gets full metadata of a product
func (a *MetadataRepositoryAdaptor) GetProductMeta(
	ctx context.Context,
	tx transactionManager.TransactionExtractor,
	id string) (entity.ProductMeta, error) {

	q := rdbms.NewQueries(a.db).WithTx(tx.TransactionPsql())
	result, err := q.GetProductMeta(ctx, id)

	if err != nil {
		return entity.ProductMeta{}, err
	}

	return entity.ProductMeta{
		Id:       result.ID,
		Name:     result.Name,
		Code:     result.Symbol,
		Exchange: result.Exchange,
		Location: result.Location.String,
		Type:     result.Type,
	}, nil
}

// This method stores metadata to metadata repository which can be mysql, radius so on...
func (a *MetadataRepositoryAdaptor) StoreProductMeta(
	ctx context.Context,
	tx transactionManager.TransactionExtractor,
	meta entity.ProductMeta) error {

	q := rdbms.NewQueries(a.db).WithTx(tx.TransactionPsql())

	return q.InsertNewProductMeta(
		ctx,
		rdbms.InsertNewProductMetaParams{
			ID:          meta.Id,
			Name:        meta.Name,
			Symbol:      meta.Code,
			Description: sql.NullString{String: meta.Description, Valid: meta.Description != ""},
			Type:        meta.Type,
			Exchange:    meta.Exchange,
			Location:    sql.NullString{String: meta.Location, Valid: meta.Type == "stock"},
		})
}

func (a *MetadataRepositoryAdaptor) Close() error {
	return a.db.Close()
}

func (a *MetadataRepositoryAdaptor) Ping() error {
	return a.db.Ping()
}

// This method get list of stored product
func (a *MetadataRepositoryAdaptor) GetStoredProductList(ctx context.Context, tx transactionManager.TransactionExtractor) ([]string, error) {
	q := rdbms.NewQueries(a.db).WithTx(tx.TransactionPsql())
	return q.GetStoredProductList(ctx)
}
