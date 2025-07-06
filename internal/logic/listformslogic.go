package logic

import (
	"context"

	"form-engine-gozero/internal/svc"
	"form-engine-gozero/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListFormsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取所有表单
func NewListFormsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListFormsLogic {
	return &ListFormsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListFormsLogic) ListForms(req *types.ListFormsReq) (resp []types.FormResp, err error) {
	// todo: add your logic here and delete this line

	return
}
