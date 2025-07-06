package handler

import (
	"net/http"

	"form-engine-gozero/internal/logic"
	"form-engine-gozero/internal/svc"
	"form-engine-gozero/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 创建表单
func CreateFormHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateFormReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewCreateFormLogic(r.Context(), svcCtx)
		resp, err := l.CreateForm(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
