// Code generated by goctl. DO NOT EDIT.
// goctl 1.8.4

package handler

import (
	"net/http"

	"form-engine-gozero/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				// 创建表单
				Method:  http.MethodPost,
				Path:    "/api/forms",
				Handler: CreateFormHandler(serverCtx),
			},
			{
				// 获取所有表单
				Method:  http.MethodGet,
				Path:    "/api/forms",
				Handler: ListFormsHandler(serverCtx),
			},
			{
				// 获取表单详情
				Method:  http.MethodGet,
				Path:    "/api/forms/:id",
				Handler: GetFormByIdHandler(serverCtx),
			},
			{
				// 提交表单数据
				Method:  http.MethodPost,
				Path:    "/api/forms/:id/entries",
				Handler: SubmitEntryHandler(serverCtx),
			},
			{
				// 分页查询表单提交数据
				Method:  http.MethodGet,
				Path:    "/api/forms/:id/entries",
				Handler: ListEntriesHandler(serverCtx),
			},
			{
				// 为指定表单添加字段
				Method:  http.MethodPost,
				Path:    "/api/forms/:id/fields",
				Handler: CreateFieldHandler(serverCtx),
			},
			{
				// 获取指定表单的字段
				Method:  http.MethodGet,
				Path:    "/api/forms/:id/fields",
				Handler: ListFieldsHandler(serverCtx),
			},
		},
	)
}
