package logic

import (
	"context"

	"form-engine-gozero/internal/svc"
	"form-engine-gozero/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListFieldsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取指定表单的字段
func NewListFieldsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListFieldsLogic {
	return &ListFieldsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListFieldsLogic) ListFields(req *types.ListFieldsReq) (resp []types.FieldResp, err error) {
	// todo: add your logic here and delete this line

	return
}
