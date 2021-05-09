package controllers

import (
	"context"

	"github.com/s-kmmr/sample-clean-architecture/interfaces/controllers/response"
	"github.com/s-kmmr/sample-clean-architecture/usecases"
)

// type Context interface{}

type MemberController interface {
	MemberList(ctx context.Context, ctrlctx Context)
}

type memberController struct {
	u usecases.MemberUseCase
}

func NewMemberController(_u usecases.MemberUseCase) MemberController {
	return &memberController{u: _u}
}

type Katsuo struct {
	Name string
}

func (m *memberController) MemberList(ctx context.Context, ctrlctx Context) {
	list, err := m.u.Members(ctx)
	if err != nil {
		ctrlctx.JSON(500, err)
		return
	}
	ctrlctx.JSON(200, response.NewMemberResponses(list))
}
