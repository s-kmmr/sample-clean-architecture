package usecases

import (
	"context"

	"github.com/s-kmmr/sample-clean-architecture/domain/model"
	"github.com/s-kmmr/sample-clean-architecture/domain/repository"
	"github.com/s-kmmr/sample-clean-architecture/domain/service/member"
	"golang.org/x/xerrors"
)

type MemberUseCase interface {
	Members(ctx context.Context) ([]model.Member, error)
	Create(ctx context.Context, member model.Member) error
}

type memberUseCase struct {
	mr repository.MemberRepository
	tr repository.TransactionRepository
	mv member.MemberNameValidator
}

func NewMemberUseCase(_mr repository.MemberRepository, _tr repository.TransactionRepository, _mv member.MemberNameValidator) MemberUseCase {
	return &memberUseCase{
		mr: _mr,
		tr: _tr,
		mv: _mv,
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
	if err := m.mv.ValidateDuplicatedName(ctx, member.LastName(), member.FirstName()); err != nil {
		return xerrors.Errorf("failed to MemberNameValidator.ValidateDuplicatedName(): %w", err)
	}
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
