package handler

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"goservertemplate/app/user/api/internal/logic"
	"goservertemplate/app/user/api/internal/svc"
	"goservertemplate/app/user/api/internal/types"
	"goservertemplate/common/response"
	"net/http"
)

func LoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Login
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewLoginLogic(r.Context(), svcCtx)
		resp, err := l.Login(&req)
		response.Response(w, resp, err)

	}
}
