package models

type UserAuth struct {
	UserId       int32       `json:"user_id"`
	RoleName     string      `json:"role_name"`
	UserIdentify string      `json:"user_identify"` // `UserIdentify` is name or nickname
	UserRoles    []*UserRole `json:"user_roles"`    // users can have many roles
}
