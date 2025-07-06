package logic

import (
	"context"

	"form-engine-gozero/internal/svc"
	"form-engine-gozero/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFormByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取表单详情
func NewGetFormByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFormByIdLogic {
	return &GetFormByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetFormByIdLogic) GetFormById(req *types.GetFormByIdReq) (resp *types.FormResp, err error) {
	// todo: add your logic here and delete this line

	return
}
