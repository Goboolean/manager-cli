package session

import "context"

func (s *Session) GetId() int {
	return s.id
}

func (s *Session) GetContext() context.Context {
	return s.ctx
}
