package res

type ErrorCode int

const (
	SettingError      ErrorCode = 1001
	AugmentError      ErrorCode = 1002
	WriteSettingError ErrorCode = 1003
)

var ErrorMap = map[ErrorCode]string{
	SettingError:      "设置错误",
	AugmentError:      "参数错误",
	WriteSettingError: "写入设置错误",
}
