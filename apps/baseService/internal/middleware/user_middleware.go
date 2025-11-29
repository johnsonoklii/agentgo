package middleware

import (
	"context"
	"github.com/go-kratos/kratos/v2/middleware"
	khttp "github.com/go-kratos/kratos/v2/transport/http"
	"github.com/johnsonoklii/agentgo/apps/baseService/internal/pkg/errorx/code"
	"github.com/johnsonoklii/agentgo/pkg/jwt"
)

func UserMiddleware(h middleware.Handler) middleware.Handler {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		r, ok := khttp.RequestFromServerContext(ctx)
		if !ok {
			return nil, code.ErrUserUnKnown
		}

		if IsWhiteList(r.URL.Path) {
			return h(ctx, req)
		}

		//reqID := r.Header.Get("X-Request-ID")
		uid := r.Header.Get("X-User-ID")
		if uid == "" {
			return nil, code.ErrUserHeader
		}

		ctx = jwt.SetUserID(ctx, uid)
		//ctx = jwt.SetReqID(ctx, reqID)

		return h(ctx, req)
	}
}
