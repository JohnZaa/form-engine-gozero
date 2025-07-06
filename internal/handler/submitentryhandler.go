package handler

import (
	"net/http"

	"form-engine-gozero/internal/logic"
	"form-engine-gozero/internal/svc"
	"form-engine-gozero/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 提交表单数据
func SubmitEntryHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SubmitEntryReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewSubmitEntryLogic(r.Context(), svcCtx)
		resp, err := l.SubmitEntry(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
