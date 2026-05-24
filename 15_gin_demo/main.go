package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 用户模型
type User struct {
	gorm.Model
	Username string `gorm:"unique;not null" json:"username"`
	Password string `gorm:"not null" json:"password"`
}

var db *gorm.DB

const secretKey = "my-secret-2026"

// 初始化数据库
func initDB() {
	dsn := "root:@tcp(127.0.0.1:3306)/ginuser?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("数据库连接失败：" + err.Error())
	}
	db.AutoMigrate(&User{})
}

// 生成JWT令牌
func generateToken(username string) (string, error) {
	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(2 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretKey))
}

// JWT鉴权中间件
func jwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.GetHeader("Authorization")
		if tokenStr == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"msg": "未登录"})
			c.Abort()
			return
		}
		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			return []byte(secretKey), nil
		})
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"msg": "令牌无效"})
			c.Abort()
			return
		}
		c.Next()
	}
}

func main() {
	initDB()
	r := gin.Default()

	// 公开接口
	r.POST("/register", func(c *gin.Context) {
		var user User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": "参数错误"})
			return
		}
		hashPwd, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"msg": "加密失败"})
			return
		}
		user.Password = string(hashPwd)
		if err := db.Create(&user).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": "用户名已存在"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"msg": "注册成功"})
	})

	r.POST("/login", func(c *gin.Context) {
		var req User
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": "参数错误"})
			return
		}
		var user User
		if err := db.Where("username=?", req.Username).First(&user).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"msg": "账号不存在"})
			return
		}
		err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"msg": "密码错误"})
			return
		}
		token, _ := generateToken(user.Username)
		c.JSON(http.StatusOK, gin.H{
			"msg":   "登录成功",
			"token": token,
		})
	})

	// 受保护接口 - 用户管理
	authGroup := r.Group("/api/users")
	authGroup.Use(jwtAuth())
	{
		// 1. 查询所有用户
		authGroup.GET("", func(c *gin.Context) {
			var users []User
			db.Find(&users)
			// 不返回密码
			type UserVO struct {
				ID        uint   `json:"id"`
				Username  string `json:"username"`
				CreatedAt string `json:"created_at"`
			}
			var userList []UserVO
			for _, u := range users {
				userList = append(userList, UserVO{
					ID:        u.ID,
					Username:  u.Username,
					CreatedAt: u.CreatedAt.Format("2006-01-02 15:04:05"),
				})
			}
			c.JSON(http.StatusOK, gin.H{
				"msg":  "查询成功",
				"data": userList,
			})
		})

		// 2. 查询单个用户（根据ID）
		authGroup.GET("/:id", func(c *gin.Context) {
			idStr := c.Param("id")
			id, err := strconv.Atoi(idStr)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"msg": "参数错误"})
				return
			}
			var user User
			if err := db.First(&user, id).Error; err != nil {
				c.JSON(http.StatusNotFound, gin.H{"msg": "用户不存在"})
				return
			}
			c.JSON(http.StatusOK, gin.H{
				"msg": "查询成功",
				"data": gin.H{
					"id":       user.ID,
					"username": user.Username,
				},
			})
		})

		// 3. 更新用户信息（用户名）
		authGroup.PUT("/:id", func(c *gin.Context) {
			idStr := c.Param("id")
			id, err := strconv.Atoi(idStr)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"msg": "参数错误"})
				return
			}
			var req struct {
				Username string `json:"username"`
			}
			if err := c.ShouldBindJSON(&req); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"msg": "参数错误"})
				return
			}
			if err := db.Model(&User{}).Where("id = ?", id).Update("username", req.Username).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"msg": "更新失败"})
				return
			}
			c.JSON(http.StatusOK, gin.H{"msg": "更新成功"})
		})

		// 4. 删除用户
		authGroup.DELETE("/:id", func(c *gin.Context) {
			idStr := c.Param("id")
			id, err := strconv.Atoi(idStr)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"msg": "参数错误"})
				return
			}
			if err := db.Delete(&User{}, id).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"msg": "删除失败"})
				return
			}
			c.JSON(http.StatusOK, gin.H{"msg": "删除成功"})
		})
	}

	r.Run("127.0.0.1:8080")
}
