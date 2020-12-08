package routers

import (
	"ExportData/src/service/impl"
	"log"

	"github.com/facebookgo/inject"
	"github.com/gin-gonic/gin"

	"ExportData/src/common/datasource"
	"ExportData/src/common/logger"
	"ExportData/src/common/middleware/cors"
	"ExportData/src/common/middleware/jwt"
	"ExportData/src/common/setting"
	"ExportData/src/controller"
	"ExportData/src/repository"
)

//InitRouter 初始化Router
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(cors.CorsHandler())
	r.Use(gin.Recovery())
	gin.SetMode(setting.Config.APP.RunMode)
	Configure(r)
	return r
}

//Configure 配置router
func Configure(r *gin.Engine) {
	//controller declare
	var user controller.UserController
	var task controller.TaskController
	//var tag cv1.Tag
	var myjwt jwt.JWT
	//inject declare
	db := datasource.Db{}
	zap := logger.Logger{}
	//Injection
	var injector inject.Graph
	if err := injector.Provide(
		&inject.Object{Value: &db},
		&inject.Object{Value: &zap},
		//&inject.Object{Value: &repository.ArticleRepository{}},
		//&inject.Object{Value: &service.ArticleService{}},
		&inject.Object{Value: &user},
		&inject.Object{Value: &task},
		&inject.Object{Value: &repository.UserRepository{}},
		&inject.Object{Value: &impl.UserService{}},
		&inject.Object{Value: &repository.UserRoleRepository{}},
		&inject.Object{Value: &impl.UserRoleService{}},
		&inject.Object{Value: &repository.TaskRepository{}},
		&inject.Object{Value: &impl.TaskService{}},
		&inject.Object{Value: &myjwt},
		&inject.Object{Value: &repository.BaseRepository{}},
	); err != nil {
		log.Fatal("inject fatal: ", err)
	}
	if err := injector.Populate(); err != nil {
		log.Fatal("injector fatal: ", err)
	}
	//zap log init
	zap.Init()
	//database connect
	if err := db.Connect(); err != nil {
		log.Fatal("db fatal:", err)
	}
	var authMiddleware = myjwt.GinJWTMiddlewareInit(&jwt.AllUserAuthorizator{})
	r.NoRoute(authMiddleware.MiddlewareFunc(), jwt.NoRouteHandler)
	r.POST("/login", authMiddleware.LoginHandler)
	r.POST("/register", user.Register)
	// 用户信息
	userAPI := r.Group("/user")
	{
		userAPI.GET("/refresh_token", authMiddleware.RefreshHandler)
	}
	// 任务服务
	taskAPI := r.Group("/api/v1/task")
	{
		taskAPI.POST("/create", task.Add)
		taskAPI.GET("/get", task.Get)
	}

	userAPI.Use(authMiddleware.MiddlewareFunc())
	{
		//userAPI.GET("/list/list", article.GetTables)
		userAPI.GET("/info", user.Info)
		userAPI.POST("/logout", user.Logout)
	}

	//var adminMiddleware = myjwt.GinJWTMiddlewareInit(&jwt.AdminAuthorizator{})
	//apiv1 := r.Group("/api/v1")
	////使用AdminAuthorizator中间件，只有admin权限的用户才能获取到接口
	//apiv1.Use(adminMiddleware.MiddlewareFunc())
	//{
	//	//vue获取table信息
	//	//apiv1.GET("/list/list", article.GetTables)
	//	apiv1.GET("/user/list", user.GetUsers)
	//	apiv1.POST("/user", user.AddUser)
	//	apiv1.PUT("/user", user.UpdateUser)
	//	apiv1.DELETE("/user/:id", user.DeleteUser)
	//	apiv1.GET("/article/list", article.GetArticles)
	//	apiv1.GET("/article/detail/:id", article.GetArticle)
	//	apiv1.POST("/article", article.AddArticle)
	//	// apiv1.PUT("/articles/:id", article.EditArticle)
	//	// apiv1.DELETE("/articles/:id", article.DeleteArticle)
	//}
}
