package repository

import "context"

type TransactionRepository interface {
	DoWithTx(ctx context.Context, f func(context.Context) error) error
}
