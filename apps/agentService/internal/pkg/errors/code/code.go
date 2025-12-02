package code

import "github.com/go-kratos/kratos/v2/errors"

var (
	ErrAgentUnKnown          = errors.New(300000, "", "internal error")
	ErrAgentNoAuth           = errors.New(300001, "", "modal no auth")
	ErrAgentNotFound         = errors.New(300002, "", "modal not found")
	ErrAgentVersionExists    = errors.New(300003, "", "agent version already exists")
	ErrAgentVersionNotFound  = errors.New(300004, "", "agent version not found")
)
