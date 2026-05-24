package controller

import (
	"gin_admin/dao"
	"gin_admin/service"
	"gin_admin/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Web 注册
func WebRegister(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	_ = service.Register(username, password)
	c.Redirect(302, "/login")
}

// Web 登录
func WebLogin(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	_, ok := service.Login(username, password)
	if !ok {
		c.String(200, "登录失败")
		return
	}
	token, _ := utils.GenerateToken(username)
	c.SetCookie("token", token, 3600, "/", "", false, false)
	c.Redirect(302, "/user/list")
}

// Web 添加用户
func WebAddUser(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	_ = service.Register(username, password)
	c.Redirect(302, "/user/list")
}

// Web 更新用户
func WebUpdateUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.PostForm("id"))
	username := c.PostForm("username")
	_ = dao.UpdateUsername(uint(id), username)
	c.Redirect(302, "/user/list")
}

// Web 删除用户
func WebDeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	_ = dao.DeleteUser(uint(id))
	c.Redirect(302, "/user/list")
}
