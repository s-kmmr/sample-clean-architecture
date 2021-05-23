package database

import "context"

type TransactionHandler interface {
	DoWithTx(ctx context.Context, f func(ctx context.Context) error) error
}
