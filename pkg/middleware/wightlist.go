package middleware

var whitelist = map[string]bool{
	"/v1/user/register": true,
	"/v1/auth/login":    true,
}

func IsWhiteList(path string) bool {
	// 一级匹配
	if whitelist[path] {
		return true
	}

	return false
}
