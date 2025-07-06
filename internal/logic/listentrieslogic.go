package logic

import (
	"context"

	"form-engine-gozero/internal/svc"
	"form-engine-gozero/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListEntriesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 分页查询表单提交数据
func NewListEntriesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListEntriesLogic {
	return &ListEntriesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListEntriesLogic) ListEntries(req *types.ListEntriesReq) (resp *types.EntryListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
