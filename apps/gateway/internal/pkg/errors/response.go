package errors

import (
	"encoding/json"
	"github.com/johnsonoklii/agentgo/apps/gateway/internal/pkg/errors/code"
	"net/http"
)

type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	//RequestID string `json:"request_id"`
	Data any `json:"data,omitempty"`
}

// 错误响应
func RespErrorCode(w http.ResponseWriter, code2 int, requestID string) {
	msg, ok := code.ErrMsg[code2]
	if !ok {
		msg = "unknown error"
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(Response{
		Code: code2,
		Msg:  msg,
		//RequestID: requestID,
	})
}

// 成功响应
func RespSuccess(w http.ResponseWriter, data any, requestID string) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(Response{
		Code: 0,
		Msg:  "success",
		//RequestID: requestID,
		Data: data,
	})
}
