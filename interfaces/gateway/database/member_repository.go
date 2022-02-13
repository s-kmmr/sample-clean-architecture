package database

import (
	"context"

	"github.com/s-kmmr/sample-clean-architecture/domain/model"
	"github.com/s-kmmr/sample-clean-architecture/domain/repository"
	"golang.org/x/xerrors"
)

type memberRepository struct {
	h MemberHandler
}

func NewMemberRepository(m MemberHandler) repository.MemberRepository {
	return &memberRepository{h: m}
}

func (r *memberRepository) FindAll() ([]model.Member, error) {
	rec, err := r.h.FindAll()
	if len(rec) < 1 {
		return nil, model.NewApiError(uint(model.MemberNotFound), xerrors.New("Member Not Found"))
	}
	if err != nil {
		return nil, xerrors.Errorf("failed to MemberHandler.FindAll(): %w", err)
	}
	return rec.MakeMembers(), nil
}

func (r *memberRepository) FindByLastFirstName(ctx context.Context, lastName string, firstName string) (*model.Member, error) {
	rec, err := r.h.FindByLastFirstName(ctx, lastName, firstName)
	if err != nil {
		return nil, xerrors.Errorf("failed to MemberHandler.FindByLastFirstName(): %w", err)
	}
	if rec == nil {
		return nil, nil
	}
	m := rec.MakeMember()
	return &m, nil
}

func (r *memberRepository) Create(ctx context.Context, member model.Member) error {
	if err := r.h.Create(ctx, member); err != nil {
		return model.NewApiError(uint(model.FailCreateMember), xerrors.Errorf("failed to MemberHandler.Create() in gateway: %w", err))
	}
	return nil
}
