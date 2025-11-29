package code

// Gateway 错误码 100000~209999
const (
	// Token 相关
	ErrTokenMissing     = 100001
	ErrTokenInvalid     = 100002
	ErrTokenLoggedOut   = 100003
	ErrUserContextEmpty = 100004
	ErrTokenDeleteFail  = 100005

	ErrRedirectMain = 100006

	// 请求参数相关
	ErrInvalidParam = 100010
	ErrUnauthorized = 100011
	ErrForbidden    = 100012

	// 服务异常
	ErrInternal = 100100
	ErrTimeout  = 100101
)

func GateInit(errMsg map[int]string) {
	errMsg[ErrTokenMissing] = "token missing"
	errMsg[ErrTokenInvalid] = "token invalid"
	errMsg[ErrTokenLoggedOut] = "token has been logged out"
	errMsg[ErrUserContextEmpty] = "user not found in context"
	errMsg[ErrTokenDeleteFail] = "failed to delete token"
	errMsg[ErrInvalidParam] = "invalid parameter"
	errMsg[ErrUnauthorized] = "unauthorized"
	errMsg[ErrForbidden] = "forbidden"
	errMsg[ErrInternal] = "internal server error"
	errMsg[ErrTimeout] = "request timeout"
	errMsg[ErrRedirectMain] = "/"
}
