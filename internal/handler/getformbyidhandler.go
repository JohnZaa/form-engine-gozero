package handler

import (
	"net/http"

	"form-engine-gozero/internal/logic"
	"form-engine-gozero/internal/svc"
	"form-engine-gozero/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 获取表单详情
func GetFormByIdHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetFormByIdReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewGetFormByIdLogic(r.Context(), svcCtx)
		resp, err := l.GetFormById(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
