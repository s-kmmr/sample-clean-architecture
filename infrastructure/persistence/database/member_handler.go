package database

import (
	"context"

	"github.com/s-kmmr/sample-clean-architecture/domain/model"
	"github.com/s-kmmr/sample-clean-architecture/infrastructure/client"
	ie "github.com/s-kmmr/sample-clean-architecture/infrastructure/persistence/entity"
	"github.com/s-kmmr/sample-clean-architecture/interfaces/gateway/database"
	ge "github.com/s-kmmr/sample-clean-architecture/interfaces/gateway/database/entity"
	"golang.org/x/xerrors"
)

type memberHandler struct {
	conn *client.SqlHandler
}

func NewMemberHandler(s *client.SqlHandler) database.MemberHandler {
	return &memberHandler{conn: s}
}

func (m *memberHandler) FindAll() (ge.MemberEntitys, error) {
	e := ie.MemberEntitys{}
	if err := m.conn.Conn().Find(&e).Error; err != nil {
		return nil, err
	}

	if len(e) < 1 {
		return nil, xerrors.New("members record is empty")
	}

	return e.MakeMembers(), nil
}

func (m *memberHandler) Create(ctx context.Context, member model.Member) error {
	tx, err := transaction(ctx)
	if err != nil {
		return xerrors.Errorf("failed to do transaction(): %w", err)
	}

	me := ie.NewMember(member)
	if err := tx.Create(&me).Error; err != nil {
		return xerrors.Errorf("failed to do tx.Create(): %w", err)
	}
	return nil
}
