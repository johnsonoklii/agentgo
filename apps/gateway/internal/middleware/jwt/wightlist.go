package jwt

var whitelist = map[string]bool{
	"/":                 true,
	"/v1/auth/login":    true,
	"/v1/user/register": true,
}

var redirectPaths = map[string]bool{
	"/v1/auth/login": true,
}

func IsWhiteList(path string) bool {
	// 一级匹配
	if whitelist[path] {
		return true
	}

	return false
}

func IsRedirect(path string) bool {
	// 一级匹配
	if redirectPaths[path] {
		return true
	}

	return false
}
