package errors

import (
	"encoding/json"
	codes "google.golang.org/grpc/codes"
	"net/http"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/transport/http/status"
)

// CustomErrorResponse 自定义错误响应
type CustomErrorResponse struct {
	Code     int         `json:"code"`
	Reason   string      `json:"reason"`
	Message  string      `json:"message"`
	MetaData interface{} `json:"details,omitempty"`
}

// CustomErrorEncoder 自定义错误编码器
func CustomErrorEncoder(w http.ResponseWriter, r *http.Request, err error) {
	w.Header().Set("Content-Type", "application/json")

	if e := errors.FromError(err); e != nil {
		w.WriteHeader(status.DefaultConverter.FromGRPCCode(codes.Code(e.Code)))
		response := CustomErrorResponse{
			Code:    int(e.Code),
			Reason:  e.Reason,
			Message: e.Message,
		}

		// 添加详细信息
		if len(e.Metadata) > 0 {
			response.MetaData = e.Metadata
		}

		json.NewEncoder(w).Encode(response)
		return
	}

	// 处理其他错误
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(CustomErrorResponse{
		Code:    http.StatusInternalServerError,
		Reason:  "INTERNAL_ERROR",
		Message: "内部服务器错误",
	})
}
