package helpers

var errCodes = map[int]string{
	10000: "参数错误",
	10001: "非法操作",
	// Book 10100 - 10199
	10100: "Book不存在",
	// 系统错误
	50000: "服务器错误，请稍后重试",
	50001: "数据获取失败，请稍后重试",
}

const (
	ErrParams       = 10000
	ErrBookNotFound = 10100
	ErrSystem       = 50000
)
