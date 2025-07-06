package handler

import (
	"net/http"

	"form-engine-gozero/internal/logic"
	"form-engine-gozero/internal/svc"
	"form-engine-gozero/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 获取指定表单的字段
func ListFieldsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ListFieldsReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewListFieldsLogic(r.Context(), svcCtx)
		resp, err := l.ListFields(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
