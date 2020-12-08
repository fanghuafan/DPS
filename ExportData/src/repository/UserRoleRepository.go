package repository

import (
	"ExportData/src/common/logger"
	"ExportData/src/models"
)

//RoleRepository 注入IDb
type UserRoleRepository struct {
	Log  logger.ILogger `inject:""`
	Base BaseRepository `inject:"inline"`
}

//GetUserRoles 获取用户身份信息
func (a *UserRoleRepository) GetUserRoles(where interface{}) []*models.UserRole {
	var roles []*models.UserRole
	if err := a.Base.Find(where, &roles, ""); err != nil {
		a.Log.Errorf("获取用户身份信息错误", err)
	}
	return roles
}

//GetRoles 获取用户角色
func (a *UserRoleRepository) GetRoles(sel *string, where interface{}) []string {
	var arrRole []string
	var roles []models.UserRole
	if err := a.Base.Find(where, &roles, *sel); err != nil {
		a.Log.Errorf("获取用户角色失败", err)
	}
	for _, role := range roles {
		arrRole = append(arrRole, role.RoleType)
	}
	return arrRole
}

//AddRole 添加用户角色
func (a *UserRoleRepository) AddRole(role *models.UserRole) bool {
	if err := a.Base.Create(&role); err != nil {
		a.Log.Errorf("添加用户角色失败", err)
		return false
	}
	return true
}

//GetRole 获取角色
func (a *UserRoleRepository) GetRole(where interface{}) *models.UserRole {
	var role models.UserRole
	if err := a.Base.First(where, &role); err != nil {
		a.Log.Errorf("获取角色失败", err)
	}
	return &role
}
