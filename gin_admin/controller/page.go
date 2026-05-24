package controller

import (
	"gin_admin/dao"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 登录页面
func LoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

// 注册页面
func RegisterPage(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", nil)
}

// 用户列表页面
func UserListPage(c *gin.Context) {
	users, _ := dao.FindAllUsers()
	c.HTML(http.StatusOK, "user_list.html", gin.H{
		"users": users,
	})
}

// 添加用户页面
func AddUserPage(c *gin.Context) {
	c.HTML(http.StatusOK, "user_form.html", gin.H{
		"user": nil,
	})
}

// 编辑用户页面
func EditUserPage(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	user, _ := dao.FindUserByID(uint(id))
	c.HTML(http.StatusOK, "user_form.html", gin.H{
		"user": user,
	})
}
