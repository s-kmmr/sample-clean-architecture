package entity

import "github.com/s-kmmr/sample-clean-architecture/domain/model"

type MemberEntity struct {
	lastName  string
	firstName string
}
type MemberEntitys []MemberEntity

func NewMemberEntity(ln string, fn string) MemberEntity {
	return MemberEntity{lastName: ln, firstName: fn}
}

func (m *MemberEntitys) MakeMembers() []model.Member {
	if len(*m) < 1 {
		return nil
	}
	members := make([]model.Member, len(*m))
	for i, v := range *m {
		members[i] = model.NewMember(v.lastName, v.firstName)
	}
	return members
}
