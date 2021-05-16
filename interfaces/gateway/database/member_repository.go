package database

import (
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

func (m *memberRepository) FindAll() ([]model.Member, error) {
	rec, err := m.h.FindAll()
	if len(rec) < 1 {
		return nil, model.NewApiError(uint(model.MemberNotFound), xerrors.New("Member Not Found"))
	}
	if err != nil {
		return nil, xerrors.Errorf("failed to FindAll(): %w", err)
	}
	return rec.MakeMembers(), nil
}
