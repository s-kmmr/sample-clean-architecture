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

// func (m *Member) SetLastName(_lastName string) {
// 	m.lastName = _lastName
// }

// func (m *Member) SetFirstName(_firstName string) {
// 	m.firstName = _firstName
// }
