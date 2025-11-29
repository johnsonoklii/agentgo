package proxy

import (
	"fmt"
	"github.com/johnsonoklii/agentgo/apps/gateway/internal/discovery"
	"github.com/johnsonoklii/agentgo/pkg/jwt"
	"io"
	"math/rand"
	"net/http"
	"time"
)

type ProxyHandler struct {
	Consul *discovery.ConsulClient
}

func NewProxyHandler(consul *discovery.ConsulClient) *ProxyHandler {
	return &ProxyHandler{Consul: consul}
}

// 转发请求到后端服务
func (p *ProxyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ProxyHandler.ServeHTTP")
	serviceName := extractServiceName(r) // 根据路径或规则提取目标服务名
	addrs, err := p.Consul.GetServiceAddresses(serviceName)
	if err != nil {
		fmt.Println("Error:", err)
		http.Error(w, "service unavailable", http.StatusServiceUnavailable)
		return
	}

	// 随机选择一个实例
	rand.Seed(time.Now().UnixNano())
	target := addrs[rand.Intn(len(addrs))]

	req, err := http.NewRequest(r.Method, "http://"+target+r.RequestURI, r.Body)
	if err != nil {
		http.Error(w, "failed to create request", http.StatusInternalServerError)
		return
	}
	req.Header = r.Header.Clone()

	userID, ok := jwt.GetUserID(r.Context())
	if ok {
		req.Header.Set("X-User-ID", userID)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, "failed to forward request", http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()

	// 返回响应
	for k, v := range resp.Header {
		w.Header()[k] = v
	}
	w.WriteHeader(resp.StatusCode)
	_, _ = io.Copy(w, resp.Body)
}

//func (p *ProxyHandler) ProxyTo(ctx context.Context, req interface{}) (interface{}, error) {
//	r, ok := khttp.RequestFromServerContext(ctx)
//	if !ok {
//		return nil, code.ErrUnKnown
//	}
//	serviceName := extractServiceName(r) // 根据路径或规则提取目标服务名
//	addrs, err := p.Consul.GetServiceAddresses(serviceName)
//	if err != nil {
//		return nil, code.ErrDiscover
//	}
//
//	fmt.Println("addrs:", addrs)
//
//	// 随机选择一个实例
//	rand.Seed(time.Now().UnixNano())
//	target := addrs[rand.Intn(len(addrs))]
//
//	// 构造请求
//	newReq, err := http.NewRequest(r.Method, "http://"+target+r.RequestURI, r.Body)
//	if err != nil {
//		return nil, code.ErrUnKnown
//	}
//	newReq.Header = r.Header.Clone()
//
//	userID, ok := jwt.GetUserID(r.Context())
//	if ok {
//		newReq.Header.Set("X-User-ID", userID)
//	}
//
//	// 执行请求
//	client := &http.Client{}
//	resp, err := client.Do(newReq)
//	if err != nil {
//		return nil, code.ErrUnKnown
//
//	}
//	//defer resp.Body.Close()
//	//
//	//// 返回响应
//	//for k, v := range resp.Header {
//	//	w.Header()[k] = v
//	//}
//	//w.WriteHeader(resp.StatusCode)
//	//_, _ = io.Copy(w, resp.Body)
//
//	return resp, nil
//}

//func Proxy_HTTP_Handler(handler *ProxyHandler) func(ctx khttp.Context) error {
//	return func(ctx khttp.Context) error {
//		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
//			return handler.ProxyTo(ctx, req)
//		})
//		out, err := h(ctx, nil)
//		if err != nil {
//			return err
//		}
//		return ctx.Result(200, out)
//	}
//}

// 提取服务名，可按规则解析路径，例如 /user/... => user-service
func extractServiceName(r *http.Request) string {
	path := r.URL.Path
	if len(path) > 1 {
		switch {
		case path == "/logout":
			return "" // 登出特殊处理，不走 Proxy
		case path[:9] == "/v1/user/" || path[:9] == "/v1/auth/":
			return "agentgo.baseService-http"
		case path[:7] == "/agent/":
			return "agent-service"
		case path[:5] == "/rag/":
			return "rag-service"
		default:
			return "default-service"
		}
	}
	return "default-service"
}
