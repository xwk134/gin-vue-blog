package res

type ErrorCode int

const (
	SettingsError ErrorCode = 1001 //系统错误
	TestError     ErrorCode = 1002 //请求超时
)

var (
	ErrorMap = map[ErrorCode]string{
		SettingsError: "系统错误",
		TestError:     "请求超时",
	}
)
