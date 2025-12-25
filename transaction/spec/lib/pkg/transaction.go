package transaction

import (
	"context"
	"io"
)

type (
	// ctxKey is the type of key used to add or retrieve a simpleTransaction from context
	ctxKey int

	// Transaction is a wrapper for a real distributed simpleTransaction selected by Route.
	Transaction interface {
		io.Closer
		EnlistResource(key string, r Resource) error
		IsMarkedRollback() bool
		Status() Status
		Commit() error
		Rollback() error
		MarkToRollback()
		GetResource(key string) Resource
	}
)

//goland:noinspection ALL
const (
	// TxKey is the value used to get/retrieve simpleTransaction from context
	TxKey ctxKey = 1
)

func FromContext(ctx context.Context) Transaction {
	if t, ok := ctx.Value(TxKey).(Transaction); ok {
		return t
	}
	return nil
}

func NewContext(parent context.Context, t Transaction) context.Context {
	return context.WithValue(parent, TxKey, t)
}
