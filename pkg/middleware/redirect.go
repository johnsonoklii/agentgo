package middleware

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	khttp "github.com/go-kratos/kratos/v2/transport/http"
)

func RedirectMiddleware(jwtManager *jwt.JWTManager, log *log.Helper) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			if r, ok := khttp.RequestFromServerContext(ctx); ok {
				if isRedirect(r.URL.Path) {
					valid, _ := isValid(ctx, jwtManager)
					if valid {
						return nil, ErrRedirectMain
					} else {
						return handler(ctx, req)
					}
				}
			}

			return handler(ctx, req)
		}
	}
}

func isRedirect(url string) bool {
	for _, path := range RedirectPaths {
		if path == url {
			return true
		}
	}

	return false
}

func isValid(ctx context.Context, jwtManager *jwt.JWTManager) (bool, error) {
	token, err := extractToken(ctx)
	if err != nil {
		log.Errorf("RedirectMiddleware.extractToken error: %v", err)
		return false, err
	}

	// 验证 Token
	_, err = jwtManager.ValidateToken(ctx, token)
	if err != nil {
		log.Errorf("RedirectMiddleware.ValidateToken error: %v", err)
		return false, err
	}

	return true, nil
}
