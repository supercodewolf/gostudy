package controller

import (
	"fmt"
	"gin_admin/dao"
	"gin_admin/service"
	"gin_admin/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 注册
func Register(c *gin.Context) {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"msg": "参数错误"})
		return
	}
	if err := service.Register(req.Username, req.Password); err != nil {
		c.JSON(400, gin.H{"msg": "注册失败"})
		return
	}
	c.JSON(200, gin.H{"msg": "注册成功"})
}

// 登录
func Login(c *gin.Context) {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"msg": "参数错误"})
		return
	}
	user, ok := service.Login(req.Username, req.Password)
	if !ok {
		c.JSON(401, gin.H{"msg": "账号或密码错误"})
		return
	}
	token, err := utils.GenerateToken(user.Username)
	if err != nil {
		fmt.Println("生成token失败：", err) // 这里会告诉你真实原因
		c.JSON(500, gin.H{"msg": "token生成失败"})
		return
	}
	fmt.Println("用户信息：", user)
	fmt.Println("生成token：", token)
	c.JSON(200, gin.H{"token": token, "msg": "登录成功"})

}

// 获取所有用户
func GetUsers(c *gin.Context) {
	users, _ := dao.FindAllUsers()
	type UserVO struct {
		ID        uint   `json:"id"`
		Username  string `json:"username"`
		CreatedAt string `json:"created_at"`
	}
	var list []UserVO
	for _, u := range users {
		list = append(list, UserVO{
			ID:        u.ID,
			Username:  u.Username,
			CreatedAt: u.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}
	c.JSON(200, gin.H{"data": list})
}

// 获取单个用户
func GetUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := dao.FindUserByID(uint(id))
	if err != nil {
		c.JSON(404, gin.H{"msg": "用户不存在"})
		return
	}
	c.JSON(200, gin.H{"data": user})
}

// 更新用户
func UpdateUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var req struct {
		Username string `json:"username"`
	}
	c.ShouldBindJSON(&req)
	dao.UpdateUsername(uint(id), req.Username)
	c.JSON(200, gin.H{"msg": "更新成功"})
}

// 删除用户
func DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := dao.DeleteUser(uint(id))
	if err != nil {
		c.JSON(404, gin.H{"msg": "用户不存在"})
		return
	}
	c.JSON(200, gin.H{"msg": "删除成功"})
}
