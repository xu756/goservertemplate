package handler

import (
	"goservertemplate/app/user/api/internal/logic"
	"goservertemplate/app/user/api/internal/svc"
	"goservertemplate/common/response"
	"net/http"
)

func InfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := logic.NewInfoLogic(r.Context(), svcCtx)
		resp, err := l.Info()
		response.Response(w, resp, err)

	}
}
