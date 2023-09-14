package productMetaRepo

import (
	"context"
	"database/sql"

	transactionManager "github.com/Goboolean/manager-cli/internal/adaptor/transaction-manager"
	"github.com/Goboolean/manager-cli/internal/domain/entity"
	"github.com/Goboolean/manager-cli/internal/infrastructure/rdbms"
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

	if !result.Location.Valid {
		result.Location.String = entity.NullString
	}

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
			Description: sql.NullString{String: meta.Description, Valid: meta.Description != entity.NullString},
			Type:        meta.Type,
			Exchange:    meta.Exchange,
			Location:    sql.NullString{String: meta.Location, Valid: meta.Location != entity.NullString},
		})
}

func (a *MetadataRepositoryAdaptor) IsProductStored(ctx context.Context, tx transactionManager.TransactionExtractor, id string) (bool, error) {
	q := rdbms.NewQueries(a.db).WithTx(tx.TransactionPsql())
	return q.CheckProductIsStored(ctx, id)
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
