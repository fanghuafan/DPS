package jwt

import (
	"encoding/json"
	"log"
	"time"

	"github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"

	"ExportData/src/common/codes"
	"ExportData/src/common/helper"
	"ExportData/src/common/logger"
	"ExportData/src/common/setting"
	"ExportData/src/models"
	"ExportData/src/service"
)

//### 如果是使用Go Module,gin-jwt模块应使用v2
//下载安装，开启Go Module "go env -w GO111MODULE=on",然后执行"go get github.com/appleboy/gin-jwt/v2"
//导入应写成 import "github.com/appleboy/gin-jwt/v2"
//### 如果不是使用Go Module
//下载安装gin-jwt，"go get github.com/appleboy/gin-jwt"
//导入import "github.com/appleboy/gin-jwt"

// JWT 注入IService
type JWT struct {
	Log       logger.ILogger           `inject:""`
	Uservice  service.IUserService     `inject:""`
	URservice service.IUserRoleService `inject:""`
}

//app 程序配置
var app = setting.Config.APP

//GinJWTMiddlewareInit 初始化中间件
func (j *JWT) GinJWTMiddlewareInit(jwtAuthorizator IAuthorizator) (authMiddleware *jwt.GinJWTMiddleware) {
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte("huyUDFHDKFJASNMDS237136478234=="),
		Timeout:     time.Minute * 30,
		MaxRefresh:  time.Hour,
		IdentityKey: app.IdentityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*models.UserAuth); ok {
				//get roles from username
				userRoles := j.URservice.GetUserRoles(v.UserId)
				jsonRole, _ := json.Marshal(userRoles)
				//maps the claims in the JWT
				return jwt.MapClaims{
					"userId":    v.UserId,
					"userName":  v.UserIdentify,
					"userRoles": helper.B2S(jsonRole),
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			roles := jwt.ExtractClaims(c)
			//extracts identity from roles
			jsonRole := roles["userRoles"].(string)
			var userRoles []*models.UserRole
			json.Unmarshal(helper.S2B(jsonRole), &userRoles)
			//Set the identity
			return &models.UserAuth{
				UserId:       int32(roles["userId"].(float64)),
				UserIdentify: roles["userName"].(string),
				UserRoles:    userRoles,
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			////handles the login logic. On success LoginResponse is called, on failure Unauthorized is called
			var loginVals models.User
			if err := c.ShouldBind(&loginVals); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			//username := c.GetString("username")
			//password := c.GetString("password")

			username := loginVals.Name
			password := loginVals.Password
			if j.Uservice.CheckUser(username, password) {
				user := j.Uservice.GetUserInfoByName(username)
				return &models.UserAuth{
					UserId:       user.Id,
					UserIdentify: username,
					UserRoles:    j.URservice.GetUserRoles(user.Id),
				}, nil
			}

			return nil, jwt.ErrFailedAuthentication
		},
		//receives identity and handles authorization logic
		Authorizator: jwtAuthorizator.HandleAuthorizator,
		//handles unauthorized logic
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		// TokenLookup is a string in the db of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		// - "param:<name>"
		TokenLookup: "header: Authorization, query: token, cookie: jwt",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	})

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}
	return
}

//NoRouteHandler 404 handler
func NoRouteHandler(c *gin.Context) {
	code := codes.PageNotFound
	c.JSON(404, gin.H{"code": code, "message": codes.GetMsg(code)})
}
