package metadataRepo

import (
	"context"
	"database/sql"

	transactionManager "github.com/Goboolean/manager-cli/internal/adaptor/transaction-manager"
	"github.com/Goboolean/manager-cli/internal/domain/entity"
	"github.com/Goboolean/shared/pkg/rdbms"
)

// This method gets unique id of a product which can be hash, UUID and so on...
func (m *MetadataRepositoryAdaptor) GetProductId(ctx context.Context,
	tx transactionManager.TransactionExtractor,
	code string) (string, error) {

	q := m.db.NewQueries().WithTx(tx.TransactionPsql())
	result, err := q.GetStockIdBySymbol(ctx, code)

	if err != nil {
		return "", err
	}

	return result, err

}

// This method gets full metadata of a product
func (m *MetadataRepositoryAdaptor) GetProductMeta(ctx context.Context,
	tx transactionManager.TransactionExtractor,
	id string) (entity.ProductMeta, error) {

	q := m.db.NewQueries().WithTx(tx.TransactionPsql())
	result, err := q.GetStockMeta(ctx, id)

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
func (m *MetadataRepositoryAdaptor) StoreProductMeta(ctx context.Context,
	tx transactionManager.TransactionExtractor,
	meta entity.ProductMeta) error {

	q := m.db.NewQueries().WithTx(tx.TransactionPsql())
	return q.InsertNewStockMeta(
		ctx,
		rdbms.InsertNewStockMetaParams{
			ID:          meta.Id,
			Name:        meta.Name,
			Symbol:      meta.Code,
			Description: sql.NullString{String: meta.Description, Valid: meta.Description != ""},
			Type:        meta.Type,
			Exchange:    meta.Exchange,
			Location:    sql.NullString{String: meta.Location, Valid: meta.Type == "stock"},
		})
}

func (m *MetadataRepositoryAdaptor) Close() error {
	return m.db.Close()
}

func (m *MetadataRepositoryAdaptor) Ping() error {
	return m.db.Ping()
}
