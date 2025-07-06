package handler

import (
	"net/http"

	"form-engine-gozero/internal/logic"
	"form-engine-gozero/internal/svc"
	"form-engine-gozero/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 分页查询表单提交数据
func ListEntriesHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ListEntriesReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewListEntriesLogic(r.Context(), svcCtx)
		resp, err := l.ListEntries(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
