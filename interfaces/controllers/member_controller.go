package controllers

import (
	"context"

	"github.com/s-kmmr/sample-clean-architecture/domain/model"
	"github.com/s-kmmr/sample-clean-architecture/interfaces/controllers/response"
	"github.com/s-kmmr/sample-clean-architecture/usecases"
	"golang.org/x/xerrors"
)

type MemberController interface {
	MemberList(ctx context.Context, ctrlctx Context)
}

type memberController struct {
	u usecases.MemberUseCase
}

func NewMemberController(_u usecases.MemberUseCase) MemberController {
	return &memberController{u: _u}
}

// @Summary 会員情報一覧をリストで返す
// @Tags member
// @Accept json
// @Produce json
// @Success 200 {object} []response.MemberResponse
// @Failure 400 {object} response.ErrorResponse
// @Router /v1/members [get]
func (m *memberController) MemberList(ctx context.Context, ctrlctx Context) {
	list, err := m.u.Members(ctx)
	if err != nil {
		var apierr *model.ApiError
		if xerrors.As(err, &apierr) {
			ctrlctx.JSON(400, response.NewErrorResponse(apierr.Code(), apierr.Err().Error()))
			return
		}
		ctrlctx.JSON(400, response.NewErrorResponse(uint(model.Unknown), err.Error()))
		return
	}
	ctrlctx.JSON(200, response.NewMemberResponses(list))

}
