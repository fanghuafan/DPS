package controller

import (
	"ExportData/src/common/codes"
	"ExportData/src/common/logger"
	"ExportData/src/models"
	"ExportData/src/service"
	"github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {
	Log     logger.ILogger       `inject:""`
	Service service.IUserService `inject:""`
}

// user register
func (u *UserController) Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBind(&user); err != nil {
		RespFail(c, http.StatusOK, codes.ErrNotExistTag, "Params null")
	}
	userName := user.Name
	password := user.Password
	u.Log.Info("query params:", userName, password)
	// check user
	isUser := u.Service.CheckUser(userName, password)
	if (isUser) {
		u.Log.Error("User already exsits!")
		RespFail(c, http.StatusOK, codes.ErrExistTag, "User already exsits")
		return
	}
	rSucc := u.Service.Register(userName, password);
	if (rSucc) {
		RespOk(c, http.StatusOK, codes.SUCCESS)
		return
	}
	RespFail(c, http.StatusOK, codes.ERROR, "User register failure")
}

func (u *UserController) Info(c *gin.Context) {
	roles := jwt.ExtractClaims(c)
	userId := int(roles["userId"].(float64))
	user := u.Service.GetUserInfo(userId);
	RespData(c, http.StatusOK, codes.SUCCESS, user)
}

//Logout 退出登录
func (a *UserController) Logout(c *gin.Context) {
	RespOk(c, http.StatusOK, codes.SUCCESS)
}
