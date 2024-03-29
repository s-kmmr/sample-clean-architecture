package database

import (
	"context"

	"github.com/s-kmmr/sample-clean-architecture/domain/model"
	"github.com/s-kmmr/sample-clean-architecture/interfaces/gateway/database/entity"
)

type MemberHandler interface {
	FindAll() (entity.MemberEntitys, error)
	FindByLastFirstName(ctx context.Context, lastName string, firstName string) (*entity.MemberEntity, error)
	Create(ctx context.Context, member model.Member) error
}
