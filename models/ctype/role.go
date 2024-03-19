package ctype

import "encoding/json"

type RoleType int

const (
	PermissionAdmin    RoleType = 1
	PermissionUser     RoleType = 2
	PermissionVisitor  RoleType = 3
	PermissionDisabled RoleType = 4
)

func (role RoleType) MarshalJSON() ([]byte, error) {
	return json.Marshal(role.String())
}

func (role RoleType) String() string {
	switch role {
	case PermissionAdmin:
		return "管理员"
	case PermissionUser:
		return "用户"
	case PermissionVisitor:
		return "访客"
	case PermissionDisabled:
		return "禁用用户"
	default:
		return "未知角色"
	}
}
