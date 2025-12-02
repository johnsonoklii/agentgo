package code

import "github.com/go-kratos/kratos/v2/errors"

var (
	ErrReq        = errors.New(000001, "", "req invalid")
	ErrUserHeader = errors.New(000002, "", "user header error")
)
