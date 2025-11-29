package jwt

import (
	"context"
	"github.com/google/uuid"
	"github.com/johnsonoklii/agentgo/apps/gateway/internal/data"
	"github.com/johnsonoklii/agentgo/apps/gateway/internal/pkg/errors"
	"github.com/johnsonoklii/agentgo/apps/gateway/internal/pkg/errors/code"
	"github.com/johnsonoklii/agentgo/pkg/jwt"
	"net/http"
)

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestID := uuid.NewString()

		if IsWhiteList(r.URL.Path) {
			next.ServeHTTP(w, r)
			return
		}

		token, err := jwt.ExtractToken(r)
		if err != nil {
			errors.RespErrorCode(w, code.ErrTokenMissing, requestID)
			return
		}

		claims, err := jwt.ParseToken(token)
		if err != nil {
			errors.RespErrorCode(w, code.ErrTokenInvalid, requestID)
			return
		}

		ok, _ := CheckToken(r.Context(), claims.UID, token)
		if !ok {
			errors.RespErrorCode(w, code.ErrTokenInvalid, requestID)
			return
		}

		ctx := jwt.SetUserID(r.Context(), claims.UID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func CheckToken(ctx context.Context, UID, token string) (bool, error) {
	key := jwt.GetRDSKeyToken(UID)
	val, err := data.Get().RDB.Get(ctx, key).Result()
	if err != nil {
		return false, nil
	}
	return val == token, nil
}

//func JWTMiddleware(h middleware.Handler) middleware.Handler {
//	return func(ctx context.Context, req interface{}) (interface{}, error) {
//		fmt.Println("xxxxxxx")
//		r, ok := khttp.RequestFromServerContext(ctx)
//		if !ok {
//			return nil, code.ErrUnKnown
//		}
//		//requestID := uuid.NewString()
//
//		if IsWhiteList(r.URL.Path) {
//			return h(ctx, req)
//		}
//
//		authHeader := r.Header.Get("Authorization")
//		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
//			return nil, code.ErrTokenInvalid
//		}
//
//		token := strings.TrimPrefix(authHeader, "Bearer ")
//		if token == "" {
//			return nil, code.ErrTokenInvalid
//		}
//
//		claims, err := jwt.ParseToken(token)
//		if err != nil {
//			return nil, code.ErrTokenInvalid
//		}
//
//		ok, _ = jwt.CheckToken(claims.UserID, token)
//		if !ok {
//			return nil, code.ErrTokenInvalid
//		}
//
//		ctx = jwt.SetUserID(ctx, claims.UserID)
//		return h(ctx, req)
//	}
//}
