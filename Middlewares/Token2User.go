package Middlewares

import (
	"dousheng-backend/Databases"
	"dousheng-backend/Databases/DAO"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Token2User() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := Databases.DatabaseSession()

		//s1 获取当前用户token
		token := c.Query("token")
		if token == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"StatusCode": 1,
				"StatusMsg":  "没有相应的token, 请重新登陆获取",
			})
			c.Abort()
			return
		}
		//s2 解析token获得用户ID var
		userID, err := ParseToken(c, token)
		if err != nil {
			c.Abort()
			return
		}
		//s3 通过IDvar获取用户表ID行记录 User model
		user, err := DAO.GetUserByID(db, userID)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"StatusCode": 3,
				"StatusMsg":  "user table 交互异常，无法获得用户ID",
			})
			c.Abort()
			return
		}
		//s4 c.set传递数据 c.next()调用后续处理函数
		c.Set("User", user)
		c.Next()
	}
}
