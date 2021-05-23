package database

import (
	"context"

	"github.com/s-kmmr/sample-clean-architecture/infrastructure/client"
	"github.com/s-kmmr/sample-clean-architecture/interfaces/gateway/database"
	"golang.org/x/xerrors"
	"gorm.io/gorm"
)

type txKey string

const contextTxKey txKey = "key"

type transactionHandler struct {
	conn *client.SqlHandler
}

func NewTransactionHandler(s *client.SqlHandler) database.TransactionHandler {
	return &transactionHandler{conn: s}
}

func (t *transactionHandler) DoWithTx(ctx context.Context, f func(ctx context.Context) error) error {
	db := t.conn.Conn()

	err := db.Transaction(func(tx *gorm.DB) error {
		ctx = context.WithValue(ctx, contextTxKey, tx)

		if err := f(ctx); err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return xerrors.Errorf("failed to Transaction: %w", err)
	}

	return nil
}

func transaction(ctx context.Context) (*gorm.DB, error) {
	v, ok := ctx.Value(contextTxKey).(*gorm.DB)
	if !ok {
		return nil, xerrors.New("failed to get transaction(begin)")
	}
	return v, nil
}
