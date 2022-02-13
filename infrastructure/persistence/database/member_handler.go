package database

import (
	"context"
	"database/sql"

	"github.com/s-kmmr/sample-clean-architecture/domain/model"
	"github.com/s-kmmr/sample-clean-architecture/infrastructure/client"
	ie "github.com/s-kmmr/sample-clean-architecture/infrastructure/persistence/entity"
	"github.com/s-kmmr/sample-clean-architecture/interfaces/gateway/database"
	ge "github.com/s-kmmr/sample-clean-architecture/interfaces/gateway/database/entity"
	"golang.org/x/xerrors"
	"gorm.io/gorm"
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

func (m *memberHandler) FindByLastFirstName(ctx context.Context, lastName string, firstName string) (*ge.MemberEntity, error) {
	e := ie.MemberEntity{}

	err := m.conn.Conn().Where("last_name = @ln AND first_name = @fn", sql.Named("ln", lastName), sql.Named("fn", firstName)).First(&e).Error
	if xerrors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, xerrors.Errorf("select error. : %w", err)
	}
	gm := e.MakeMember()
	return &gm, nil
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
