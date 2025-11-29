package middleware

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	khttp "github.com/go-kratos/kratos/v2/transport/http"
	"strings"
)

// AuthMiddleware JWT 认证中间件
func AuthMiddleware(jwtManager *jwt.JWTManager, log *log.Helper) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			// 检查是否是免认证路径
			if httpReq, ok := khttp.RequestFromServerContext(ctx); ok {
				if isWhitelist(httpReq.URL.Path) {
					return handler(ctx, req)
				}
			}

			token, err := extractToken(ctx)
			if err != nil {
				log.Errorf("AuthMiddleware.extractToken error: %v", err)
				return nil, err
			}

			// 验证 Token
			claims, err := jwtManager.ValidateToken(ctx, token)
			if err != nil {
				log.Errorf("AuthMiddleware.ValidateToken error: %v", err)
				return nil, err
			}

			// 将用户信息存入上下文
			ctx = withAuthContext(ctx, claims)

			return handler(ctx, req)
		}
	}
}

// extractToken 从请求头提取 Token
func extractToken(ctx context.Context) (string, error) {
	if req, ok := khttp.RequestFromServerContext(ctx); ok {
		authHeader := req.Header.Get("Authorization")
		if authHeader == "" {
			return "", ErrNoAuth
		}

		// Bearer Token 格式
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			return "", ErrTokenInvalid
		}

		return parts[1], nil
	}
	return "", ErrTokenInvalid
}

// isWhitelist 检查路径是否在免认证列表中
func isWhitelist(path string) bool {
	for _, pattern := range DefaultWhitelist {
		if pattern == path {
			return true
		}
		if strings.HasSuffix(pattern, "/") && strings.HasPrefix(path, pattern) {
			return true
		}
	}
	return false
}
