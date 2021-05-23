package controllers

import (
	"context"

	"github.com/s-kmmr/sample-clean-architecture/domain/model"
	"github.com/s-kmmr/sample-clean-architecture/interfaces/controllers/request"
	"github.com/s-kmmr/sample-clean-architecture/interfaces/controllers/response"
	"github.com/s-kmmr/sample-clean-architecture/usecases"
	"golang.org/x/xerrors"
)

type MemberController interface {
	MemberList(ctx context.Context, ctrlctx Context)
	CreateMember(ctx context.Context, ctrlctx Context)
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

// @Summary 会員情報の登録
// @Tags member
// @Accept json
// @Produce json
// @Param member body request.MemberRequest true "Add New Member"
// @Success 200 {object} response.Success
// @Failure 400 {object} response.ErrorResponse
// @Router /v1/members [post]
func (m *memberController) CreateMember(ctx context.Context, ctrlctx Context) {
	req := request.MemberRequest{}
	if err := ctrlctx.Bind(&req); err != nil {
		ctrlctx.JSON(400, response.NewErrorResponse(uint(model.Unknown), err.Error()))
		return
	}
	member := model.NewMember(req.LastName, req.FirstName)

	if err := m.u.Create(ctx, member); err != nil {
		var apierr *model.ApiError
		if xerrors.As(err, &apierr) {
			ctrlctx.JSON(400, response.NewErrorResponse(apierr.Code(), apierr.Err().Error()))
			return
		}
		ctrlctx.JSON(400, response.NewErrorResponse(uint(model.Unknown), err.Error()))
		return
	}
	ctrlctx.JSON(200, response.NewSuccess("join member!"))
}
