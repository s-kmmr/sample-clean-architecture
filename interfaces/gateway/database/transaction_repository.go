package database

import (
	"context"

	"github.com/s-kmmr/sample-clean-architecture/domain/repository"
	"golang.org/x/xerrors"
)

type transactionRepository struct {
	h TransactionHandler
}

func NewTransactionRepository(t TransactionHandler) repository.TransactionRepository {
	return &transactionRepository{h: t}
}

func (t *transactionRepository) DoWithTx(ctx context.Context, f func(context.Context) error) error {

	if err := t.h.DoWithTx(ctx, f); err != nil {
		return xerrors.Errorf("failed to DoWithTx(): %w", err)
	}

	return nil
}
