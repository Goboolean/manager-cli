package out

import (
	"context"

	"github.com/Goboolean/manager-cli/internal/domain/entity/session"
)

type TransactorPort interface {
	CreateNewSession(ctx context.Context) *session.Session
	Begin(*session.Session) error
	Commit(*session.Session) error
	Rollback(*session.Session) error
}
