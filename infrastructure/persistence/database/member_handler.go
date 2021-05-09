package database

import (
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

func (u *memberHandler) FindAll() (ge.MemberEntitys, error) {
	e := ie.MemberEntitys{}
	if err := u.conn.Conn().Find(&e).Error; err != nil {
		return nil, err
	}

	if len(e) < 1 {
		return nil, xerrors.New("members record is empty")
	}

	return e.MakeMembers(), nil
}
