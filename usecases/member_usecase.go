package usecases

import (
	"context"

	"github.com/s-kmmr/sample-clean-architecture/domain/model"
	"github.com/s-kmmr/sample-clean-architecture/domain/repository"
	"golang.org/x/xerrors"
)

type MemberUseCase interface {
	Members(ctx context.Context) ([]model.Member, error)
	Create(ctx context.Context, member model.Member) error
}

type memberUseCase struct {
	mr repository.MemberRepository
	tr repository.TransactionRepository
}

func NewMemberUseCase(_mr repository.MemberRepository, _tr repository.TransactionRepository) MemberUseCase {
	return &memberUseCase{
		mr: _mr,
		tr: _tr,
	}
}

func (m *memberUseCase) Members(ctx context.Context) ([]model.Member, error) {
	members, err := m.mr.FindAll()
	if err != nil {
		return nil, xerrors.Errorf("failed to FindAll(): %w", err)
	}
	return members, nil
}

func (m *memberUseCase) Create(ctx context.Context, member model.Member) error {
	err := m.tr.DoWithTx(ctx, func(ctx context.Context) error {
		if err := m.mr.Create(ctx, member); err != nil {
			return xerrors.Errorf("failed to Create() in transaction: %w", err)
		}
		return nil
	})

	if err != nil {
		return xerrors.Errorf("failed to DoWithTx(): %w", err)
	}
	return nil
}
