package metadataRepo

import (
	"context"
	"database/sql"

	"github.com/Goboolean/manager-cli/internal/domain/entity"
	"github.com/Goboolean/manager-cli/internal/domain/entity/session"
	"github.com/Goboolean/shared/pkg/rdbms"
)

func (m *MetadataRepositoryAdaptor) isExpiredSession(session *session.Session) bool {
	if _, ok := m.transactions[session.GetId()]; !ok {
		return true
	}

	if _, ok := m.queries[session.GetId()]; !ok {
		return true
	}

	return false

}

func (m *MetadataRepositoryAdaptor) CreateTxSession(ctx context.Context) (*session.Session, error) {

	instance := session.New(ctx)
	tx, err := m.db.NewTx(ctx)

	if err != nil {
		return nil, err
	}

	m.transactions[instance.GetId()] = tx.Transaction().(*sql.Tx)
	m.queries[instance.GetId()] = m.db.NewQueries().WithTx(tx.Transaction().(*sql.Tx))
	return instance, nil
}

func (m *MetadataRepositoryAdaptor) Commit(session *session.Session) error {

	if m.isExpiredSession(session) {
		return errExpiredSession
	}

	err := m.transactions[session.GetId()].Commit()
	if err != nil {
		return err
	}

	delete(m.transactions, session.GetId())
	delete(m.queries, session.GetId())
	return nil
}

func (m *MetadataRepositoryAdaptor) Rollback(session *session.Session) error {

	if m.isExpiredSession(session) {
		return errExpiredSession
	}

	err := m.transactions[session.GetId()].Rollback()
	if err != nil {
		return err
	}
	delete(m.transactions, session.GetId())
	delete(m.queries, session.GetId())
	return nil
}

// This method gets unique id of a product which can be hash, UUID and so on...
func (m *MetadataRepositoryAdaptor) GetProductId(session *session.Session, code string) (string, error) {

	result, err := m.queries[session.GetId()].GetStockIdBySymbol(session.GetContext(), code)

	if err != nil {
		return "", err
	}

	return result, err

}

// This method gets full metadata of a product
func (m *MetadataRepositoryAdaptor) GetProductMeta(session *session.Session, id string) (entity.ProductMeta, error) {

	result, err := m.queries[session.GetId()].GetStockMeta(session.GetContext(), id)

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
func (m *MetadataRepositoryAdaptor) StoreProductMeta(session *session.Session, meta entity.ProductMeta) error {

	return m.queries[session.GetId()].InsertNewStockMeta(
		session.GetContext(),
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
