package out

import "context"

type TransactorPort interface {
	Begin() error
	Commit() error
	Rollback() error
	Context() context.Context
}
