package repository

import (
	"context"

	"github.com/s-kmmr/sample-clean-architecture/domain/model"
)

type MemberRepository interface {
	FindAll() ([]model.Member, error)
	Create(ctx context.Context, member model.Member) error
}
