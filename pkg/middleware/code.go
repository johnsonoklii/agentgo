package middleware

import (
	"github.com/go-kratos/kratos/v2/errors"
)

var (
	ErrRedirectMain = errors.New(302, "", "/")
)

var (
	ErrTokenInvalid = jwt.ErrTokenInvalid
	ErrTokenExpired = jwt.ErrTokenExpired
	ErrTokenError   = jwt.ErrTokenError
	ErrNoAuth       = jwt.ErrNoAuth
)
