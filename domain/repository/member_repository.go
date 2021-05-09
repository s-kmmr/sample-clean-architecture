package repository

import "github.com/s-kmmr/sample-clean-architecture/domain/model"

type MemberRepository interface {
	FindAll() ([]model.Member, error)
}
