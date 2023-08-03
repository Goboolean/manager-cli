package out

import (
	"context"

	"github.com/Goboolean/manager-cli/internal/domain/entity/session"
)

type TransactorPort interface {
	// Create new transaction and begin it.
	// Returns a session object to manage each session of the transaction.
	CreateTxSession(ctx context.Context) (*session.Session, error)
	Commit(session *session.Session) error
	Rollback(session *session.Session) error
}
