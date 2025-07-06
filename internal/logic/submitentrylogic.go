package logic

import (
	"context"

	"form-engine-gozero/internal/svc"
	"form-engine-gozero/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SubmitEntryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 提交表单数据
func NewSubmitEntryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SubmitEntryLogic {
	return &SubmitEntryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SubmitEntryLogic) SubmitEntry(req *types.SubmitEntryReq) (resp *types.FormEntryResp, err error) {
	// todo: add your logic here and delete this line

	return
}
