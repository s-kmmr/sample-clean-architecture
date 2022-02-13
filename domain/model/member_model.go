package model

type Member struct {
	lastName  string
	firstName string
}

func NewMember(ln string, fn string) Member {
	return Member{lastName: ln, firstName: fn}
}

func (m Member) LastName() string {
	return m.lastName
}

func (m Member) FirstName() string {
	return m.firstName
}
