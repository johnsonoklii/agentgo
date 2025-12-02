package code

import "github.com/go-kratos/kratos/v2/errors"

var (
	ErrUserUnKnown  = errors.New(200000, "", "internal error")
	ErrUserNotFound = errors.New(200001, "", "user not exist")
	ErrUserExisted  = errors.New(200002, "", "user already exist")
	ErrUserInfo     = errors.New(200003, "", "user info error")
	ErrUserHeader   = errors.New(200004, "", "user header error")
	ErrUserNoAuth   = errors.New(200005, "", "user no auth")
)
