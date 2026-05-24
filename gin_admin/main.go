package main

import (
	"gin_admin/config"
	"gin_admin/controller"
	"gin_admin/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化数据库
	config.InitDB()

	r := gin.Default()

	// 加载 View 层
	r.LoadHTMLGlob("templates/*")

	// 页面路由
	r.GET("/login", controller.LoginPage)
	r.GET("/register", controller.RegisterPage)
	r.POST("/login", controller.WebLogin)
	r.POST("/register", controller.WebRegister)

	// 需要登录的页面
	web := r.Group("/user")
	web.Use(middleware.JWTAuthCookie()) // 新增支持Cookie登录
	{
		web.GET("/list", controller.UserListPage)
		web.GET("/add", controller.AddUserPage)
		web.POST("/add", controller.WebAddUser)
		web.GET("/edit", controller.EditUserPage)
		web.POST("/update", controller.WebUpdateUser)
		web.GET("/delete", controller.WebDeleteUser)
	}

	//公开接口
	r.POST("/register", controller.Register)
	r.POST("/login", controller.Login)

	// 需要登录的接口
	auth := r.Group("/api")
	auth.Use(middleware.JWTAuth())
	{
		auth.GET("/users", controller.GetUsers)
		auth.GET("/users/:id", controller.GetUser)
		auth.PUT("/users/:id", controller.UpdateUser)
		auth.DELETE("/users/:id", controller.DeleteUser)
	}

	// 启动
	r.Run("127.0.0.1:8080")
}
