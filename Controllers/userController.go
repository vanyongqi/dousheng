package Controllers

import (
	"dousheng-backend/Controllers/common"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"sync/atomic"
)

/*
	demo代码，注册、登录、个人信息三个函数，定义了一个map[string]User变量 做测试

// user data will be cleared every time the server starts
// test data: username=zhanglei, password=douyin
func Register(c *gin.Context) {}
func Login(c *gin.Context) {}
func UserInfo(c *gin.Context) {}
*/
var usersLoginInfo = map[string]common.User{"fanyongqi": {
	Id:            1,
	Name:          "fanyongqi",
	Password:      "123456",
	FollowCount:   10,
	FollowerCount: 5,
	IsFollow:      true,
}}
var userIdSequence = int64(1)

/**************** var block ***********************/

type RegisterRequest struct {
	Username string `json:"username" binding:"required,min=6,max=16" form:"username"`
	Password string `json:"password" binding:"required,min=6,max=16" form:"password"`
}

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

type UserLoginResponse struct {
	Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	Response
	User common.User `json:"user"`
}

func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	token := username + password

	if _, exist := usersLoginInfo[token]; exist {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User already exist"},
		})
	} else {
		atomic.AddInt64(&userIdSequence, 1)
		newUser := common.User{
			Id:   userIdSequence,
			Name: username,
		}
		usersLoginInfo[token] = newUser
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 0},
			UserId:   userIdSequence,
			Token:    username + password,
		})
	}
}

/***********************func block*********************************/

func HashEncode(str string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(str), bcrypt.DefaultCost)
	if err != nil {
		logrus.Info("调用bcrypt 加密异常，编码失败:%w", err)
	}
	return string(hash)
}

func VerifyPassword(sourcePwd, hashPwd string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashPwd), []byte(sourcePwd)) == nil
}

func PostRegisterUser(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
}
func PostLoginUser(c *gin.Context) {

}

func GetUserInfo(c *gin.Context) {

}
