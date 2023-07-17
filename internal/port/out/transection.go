package out

import "context"

type TransactorPort interface {
	Begin(ctx context.Context) error
	Commit() error
	Rollback() error
	Context() context.Context
}
