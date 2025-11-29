package middleware

import (
	"strings"
)

// WhitelistManager 免认证路径管理器
type WhitelistManager struct {
	paths []string
}

func NewWhitelistManager(paths []string) *WhitelistManager {
	return &WhitelistManager{paths: paths}
}

func (wm *WhitelistManager) AddPath(path string) {
	wm.paths = append(wm.paths, path)
}

func (wm *WhitelistManager) IsWhitelist(path string) bool {
	for _, pattern := range wm.paths {
		if pattern == path {
			return true
		}
		if strings.HasSuffix(pattern, "/*") {
			prefix := strings.TrimSuffix(pattern, "/*")
			if strings.HasPrefix(path, prefix) {
				return true
			}
		}
		if strings.HasSuffix(pattern, "/") && strings.HasPrefix(path, pattern) {
			return true
		}
	}
	return false
}

// 默认免认证路径
var DefaultWhitelist = []string{
	"/v1/auth/login",
	"/v1/user/register",
	"/v1/healthz",
}

// 需要重定向的路径（已登录用户访问这些路径会跳转首页）
var RedirectPaths = []string{
	"/v1/auth/login",
}
