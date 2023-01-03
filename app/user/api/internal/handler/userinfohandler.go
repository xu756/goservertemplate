package handler

import (
	"goservertemplate/app/user/api/internal/logic"
	"goservertemplate/app/user/api/internal/svc"
	"goservertemplate/common/response"
	"net/http"
)

func userInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := logic.NewUserInfoLogic(r.Context(), svcCtx)
		resp, err := l.UserInfo()
		response.Response(w, resp, err)

	}
}
