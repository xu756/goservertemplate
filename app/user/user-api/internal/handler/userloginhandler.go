package handler

import (
	"goservertemplate/app/user/user-api/internal/logic"
	"goservertemplate/app/user/user-api/internal/svc"
	"goservertemplate/app/user/user-api/internal/types"
	"greet/response" // ①
	"net/http"
)

func userLoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserLogin
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewUserLoginLogic(r.Context(), svcCtx)
		resp, err := l.UserLogin(&req)
		response.Response(w, resp, err) //②

	}
}
