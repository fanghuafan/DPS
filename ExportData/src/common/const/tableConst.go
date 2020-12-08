package _const

type ConstByTable int

const (
	// 超级管理员
	USER_ACTIVE ConstByTable = iota
	// 管理员
	USER_DISACTIVE
)

type RoleType string

const (
	// 超级管理员
	SUPER_ADMIN RoleType = "super_admin"
	// 管理员
	ADMIN RoleType = "admin"
	// 管理者（或主管）
	MANAGER RoleType = "manager"
	// 普通用户
	ORDINARY RoleType = "ordinary"
)


