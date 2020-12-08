package impl

import (
	"ExportData/src/common/const"
	"ExportData/src/common/logger"
	"ExportData/src/models"
	"ExportData/src/repository"
	"github.com/astaxie/beego/logs"
	"time"
)

type UserService struct {
	Repository         repository.UserRepository     `inject:"inline"`
	UserRoleRepository repository.UserRoleRepository `inject:"inline"`
	Log                logger.ILogger                `inject:""`
}

//CheckUser 身份验证
func (a *UserService) CheckUser(username string, password string) bool {
	where := models.User{Name: username, Password: password}
	return a.Repository.CheckUser(&where)
}

//GetUserInfo 获取用户信息
func (a *UserService) GetUserInfo(userId int) *models.User {
	return a.Repository.GetUserByID(userId)
}

// user register
func (u *UserService) Register(username string, password string) bool {
	user := models.User{
		Name:       username,
		Password:   password,
		CreateTime: time.Now(),
		State:      int(_const.USER_ACTIVE),
	}
	if (u.Repository.AddUser(&user)) {
		id := user.Id
		logs.Debug("User create success ok:", id)
		userRole := models.UserRole{
			Id:       id,
			RoleName: "普通用户组",
			RoleType: string(_const.ORDINARY),
		}
		if (u.UserRoleRepository.AddRole(&userRole)) {
			return true
		}
	}
	return false
}

func (u *UserService) GetUserInfoByName(userName string) *models.User {
	idName := "id"
	userId := u.Repository.GetUserID(&idName, "name = \""+userName+"\"")
	user := u.Repository.GetUserByID(int(userId))
	return user
}
