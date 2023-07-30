package session

import (
	"context"
	"math/rand"
	"time"
)

type Session struct {
	// Unique identifier of each session
	id  int
	ctx context.Context

	// conteains extra value of a session
	// If loose typing becomes a problem,
	// consider creating different structures for each usecase.
	Info map[string]interface{}
}

func New(ctx context.Context) *Session {
	rand.Seed(time.Now().Unix())

	return &Session{
		id:  rand.Int(),
		ctx: ctx,
	}
}
