package member

import (
	"context"

	"github.com/s-kmmr/sample-clean-architecture/domain/repository"
	"golang.org/x/xerrors"
)

type MemberNameValidator interface {
	ValidateDuplicatedName(ctx context.Context, lastName string, firstName string) error
}

type memberNameValidator struct {
	memberRepository repository.MemberRepository
}

func NewNameValidator(_mr repository.MemberRepository) MemberNameValidator {
	return &memberNameValidator{
		memberRepository: _mr,
	}
}

func (v *memberNameValidator) ValidateDuplicatedName(ctx context.Context, lastName string, firstName string) error {
	mem, err := v.memberRepository.FindByLastFirstName(ctx, lastName, firstName)
	if err != nil {
		return xerrors.Errorf("failed to memberRepository.FindByName(): %w", err)
	}
	if mem != nil {
		return xerrors.Errorf("lastName(%s) and firstName(%s) is duplicated", lastName, firstName)
	}
	return nil
}
