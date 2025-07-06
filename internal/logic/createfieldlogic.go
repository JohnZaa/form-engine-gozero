package logic

import (
	"context"

	"form-engine-gozero/internal/svc"
	"form-engine-gozero/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateFieldLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 为指定表单添加字段
func NewCreateFieldLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateFieldLogic {
	return &CreateFieldLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateFieldLogic) CreateField(req *types.CreateFieldReq) (resp *types.FieldResp, err error) {
	// todo: add your logic here and delete this line

	return
}
