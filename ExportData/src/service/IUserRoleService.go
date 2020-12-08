package service

import "ExportData/src/models"

type IUserRoleService interface {
	GetUserRoles(userId int32) []*models.UserRole

	// 添加用户角色
	AddUserRole(role models.UserRole) bool
}
