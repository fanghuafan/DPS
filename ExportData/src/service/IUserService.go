package service

import "ExportData/src/models"

type IUserService interface {
	// CheckUser 身份验证
	CheckUser(username string, password string) bool

	// 获取用户信息
	GetUserInfo(userId int) *models.User

	// get user by name
	GetUserInfoByName(userName string) *models.User

	// user register
	Register(username string, password string) bool
}
