package ctype

import "encoding/json"

type SignStatus int

const (
	SignWithQQ    SignStatus = 1
	SignWithEmail SignStatus = 2
)

func (s SignStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

func (s SignStatus) String() string {
	switch s {
	case SignWithQQ:
		return "QQ登录"
	case SignWithEmail:
		return "Email登录"
	default:
		return "未知"
	}
}
