package jwt

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/johnsonoklii/agentgo/apps/gateway/internal/pkg/errors"
	"github.com/johnsonoklii/agentgo/apps/gateway/internal/pkg/errors/code"
	"github.com/johnsonoklii/agentgo/pkg/jwt"

	"net/http"
)

func RedirectMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if IsRedirect(r.URL.Path) {
			valid, _ := isValid(r)
			if valid {
				errors.RespErrorCode(w, code.ErrRedirectMain, "")
				return
			} else {
				next.ServeHTTP(w, r)
				return
			}
		}

		next.ServeHTTP(w, r)
		return
	})
}

func isValid(r *http.Request) (bool, error) {
	token, err := jwt.ExtractToken(r)
	if err != nil {
		return false, err
	}

	// 验证 Token
	_, err = jwt.ParseToken(token)
	if err != nil {
		log.Errorf("RedirectMiddleware.ValidateToken error: %v", err)
		return false, err
	}

	return true, nil
}
