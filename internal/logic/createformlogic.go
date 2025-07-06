package logic

import (
	"context"

	"form-engine-gozero/internal/svc"
	"form-engine-gozero/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateFormLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 创建表单
func NewCreateFormLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateFormLogic {
	return &CreateFormLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateFormLogic) CreateForm(req *types.CreateFormReq) (resp *types.FormResp, err error) {
	// todo: add your logic here and delete this line

	return
}
