package entity

import (
	ge "github.com/s-kmmr/sample-clean-architecture/interfaces/gateway/database/entity"
)

type MemberEntity struct {
	LastName  string
	FirstName string
}
type MemberEntitys []MemberEntity

func (e *MemberEntity) TableName() string {
	return "member"
}

func (e *MemberEntitys) MakeMembers() []ge.MemberEntity {
	if len(*e) < 1 {
		return nil
	}
	members := make([]ge.MemberEntity, len(*e))
	for i, v := range *e {
		members[i] = ge.NewMemberEntity(v.LastName, v.FirstName)
	}

	return members
}
