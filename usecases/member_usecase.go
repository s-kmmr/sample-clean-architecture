package usecases

import (
	"context"

	"github.com/s-kmmr/sample-clean-architecture/domain/model"
	"github.com/s-kmmr/sample-clean-architecture/domain/repository"
	"golang.org/x/xerrors"
)

type MemberUseCase interface {
	Members(ctx context.Context) ([]model.Member, error)
}

type memberUseCase struct {
	mr repository.MemberRepository
}

func NewMemberUseCase(_mr repository.MemberRepository) MemberUseCase {
	return &memberUseCase{mr: _mr}
}

func (m *memberUseCase) Members(ctx context.Context) ([]model.Member, error) {
	members, err := m.mr.FindAll()
	if err != nil {
		return nil, xerrors.Errorf("failed to FindAll(): %w", err)
	}
	return members, nil
}
