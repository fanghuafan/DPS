package impl

import (
	"ExportData/src/models"
	"ExportData/src/repository"
)

// RoleService IRoleRepository
type UserRoleService struct {
	Repository repository.UserRoleRepository `inject:"inline"`
}

func (c *UserRoleService) GetUserRoles(userId int32) []*models.UserRole {
	where := models.UserRole{Id: userId}
	return c.Repository.GetUserRoles(&where)
}

func (c *UserRoleService) AddUserRole(role models.UserRole) bool {
	return c.Repository.AddRole(&role)
}
