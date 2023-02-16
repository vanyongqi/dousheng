package Middlewares

import (
	"dousheng-backend/Controllers/common"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

const (
	SECRETKEY = "243423ffslsfsldfl412fdsfsdf" //私钥
)

// CreateToken JWT is JSON Web Tokens ！
// 创建Token
func CreateToken(userID uint, userName string) (string, error) {
	//方式1： 标准claim
	//claims := &jwt.StandardClaims{
	//    ExpiresAt: time.Now().Add(time.Duration(maxAge)*time.Second).Unix(), // 过期时间，必须设置,
	//    Issuer:    "frankie",// 非必须，也可以填充用户名，
	//}
	claims := jwt.MapClaims{
		"id":       userID,
		"username": userName,
		"nbf":      time.Now().Unix(),
	}
	//方式2： 自定义claim
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(SECRETKEY))
	if err != nil {
		logrus.Warn("创建Token 过程异常\n")
	}
	return tokenString, err
}

// ParseToken 解析Token
func ParseToken(ctx *gin.Context, tokenString string) (uint, string, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("加密方式错误: %v", t.Header["alg"])
		}
		return []byte(SECRETKEY), nil
	})
	if err != nil {
		logrus.Warn("登录token过程出现异常\n")
		ctx.JSON(http.StatusUnauthorized, common.Response{
			StatusCode: common.REQUESTERROR,
			StatusMsg:  "token获取错误, 请重新登陆获取",
		})
		return 0, "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !(ok && token.Valid) {
		logrus.Warn("token解析过程出现异常\n")
		ctx.JSON(http.StatusUnauthorized, common.Response{
			StatusCode: common.REQUESTERROR,
			StatusMsg:  "token解析错误, 请重新登陆获取",
		})
		return 0, "", err
	}
	userID := uint(claims["id"].(float64))
	userName := claims["name"].(string)
	createTime, _ := claims["nbf"].(int64)

	if time.Now().Unix()-createTime > 36*int64(time.Hour) {
		//	logrus.Warn("token超时出现异常\n")//不算异常不需要记录。
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"StatusCode": 2,
			"StatusMsg":  "token超时, 请重新登陆获取",
		})
		return 0, "", err
	}

	return userID, userName, nil
}
