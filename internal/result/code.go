package result

// 成功
var OK Result = New(0, "success")

// 基础 10000 - 10099
var (
	ErrParams       Result = New(10000, "参数错误")
	ErrIllegal      Result = New(10001, "身份验证失败，请重新授权")
	ErrLoginExpired Result = New(10002, "登录已过期，请重新授权")
)

// 用户 10100 - 10199
var ErrUserNotFound Result = New(10100, "用户不存在")

// 系统错误
var (
	ErrSystem Result = New(50000, "服务器错误，请稍后重试")
	ErrIao    Result = New(50001, "内部服务请求失败，请稍后重试")
)
