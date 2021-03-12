package helpers

type Code int

var codeM = map[Code]string{
	10000: "参数错误",
	10001: "身份验证失败，请重新授权",
	10002: "登录已过期，请重新授权",
	// 书籍 10100 - 10199
	10100: "无效的书籍",
	// 系统错误
	50000: "服务器错误，请稍后重试",
}

const (
	ErrParams       Code = 10000
	ErrIllegal      Code = 10001
	ErrLoginExpired Code = 10002
	ErrInvalidBook  Code = 10100
	ErrSystem       Code = 50000
)
