package out

import (
	"context"

	"github.com/Goboolean/manager-cli/internal/domain/entity"
)

type TransactorPort interface {
	CreateNewSession(ctx context.Context) *entity.Session
	Begin(*entity.Session) error
	Commit(*entity.Session) error
	Rollback(*entity.Session) error
}
