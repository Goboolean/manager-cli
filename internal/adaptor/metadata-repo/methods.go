package metadataRepo

import (
	"context"
	"database/sql"

	"github.com/Goboolean/manager-cli/internal/domain/entity"
	"github.com/Goboolean/shared/pkg/rdbms"
)

func (m *MetadataRepositoryAdaptor) Begin(ctx context.Context) error {
	m.ctx = ctx

	transactor, err := m.db.NewTx(ctx)

	if err != nil {
		return err
	}

	m.tx = transactor.Transaction().(*sql.Tx)
	m.q = m.q.WithTx(m.tx)

	return nil
}

func (m *MetadataRepositoryAdaptor) Commit() error {
	return m.tx.Commit()
}

func (m *MetadataRepositoryAdaptor) Rollback() error {
	return m.tx.Rollback()
}

func (m *MetadataRepositoryAdaptor) Context() context.Context {
	return m.ctx
}

// This method gets unique id of a product which can be hash, UUID and so on...
func (m *MetadataRepositoryAdaptor) GetProductId(code string) (string, error) {

	result, err := m.q.GetStockIdBySymbol(m.ctx, code)

	if err != nil {
		return "", err
	}

	return result, err

}

// This method gets full metadata of a product
func (m *MetadataRepositoryAdaptor) GetProductMeta(id string) (entity.ProductMeta, error) {

	result, err := m.q.GetStockMeta(m.ctx, id)

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
func (m *MetadataRepositoryAdaptor) StoreProductMeta(meta entity.ProductMeta) error {

	return m.q.InsertNewStockMeta(m.ctx, rdbms.InsertNewStockMetaParams{
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
