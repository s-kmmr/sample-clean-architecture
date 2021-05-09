package database

import "github.com/s-kmmr/sample-clean-architecture/interfaces/gateway/database/entity"

type MemberHandler interface {
	FindAll() (entity.MemberEntitys, error)
}
