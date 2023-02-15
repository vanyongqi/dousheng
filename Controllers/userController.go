package Controllers

import (
	"dousheng-backend/Controllers/common"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

/*
	demo代码，注册、登录、个人信息三个函数，定义了一个map[string]User变量 做测试
TEST code start
// user data will be cleared every time the server starts
// test data: username=zhanglei, password=douyin
func Register(c *gin.Context) {}
func Login(c *gin.Context) {}
func UserInfo(c *gin.Context) {}
*/
/*
var usersLoginInfo = map[string]common.User{"fanyongqi": {
	Id:            1,
	Name:          "fanyongqi",
	Password:      "123456",
	FollowCount:   10,
	FollowerCount: 5,
	IsFollow:      true,
}}
var userIdSequence = int64(1)


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

//TEST is over.
*/

/**************** var block ***********************/

type Register_Login_Request struct {
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

/***********************func block*********************************/

func BcryptEncode(str string) string {
	bcryptPassword, err := bcrypt.GenerateFromPassword([]byte(str), bcrypt.DefaultCost)
	if err != nil {
		logrus.Info("调用bcrypt,加密异常，HASH编码失败:%w", err)
	}
	return string(bcryptPassword)
}

func VerifyPassword(sourcePwd, hashPwd string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashPwd), []byte(sourcePwd)) == nil
}

func PostRegisterUser(c *gin.Context) {
	var req Register_Login_Request
	// s1 verify parameters is valid
	// s2 verify user is already registered or not
	// s3 create user from GROM by username and password
	// s4 generate token from middleware layer
	// s5 response about user's registration
	c.JSON(200, gin.H{"user": "register is okay"})

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

}

func PostLoginUser(c *gin.Context) {

	var req Register_Login_Request
	// s1 verify parameters is valid
	// s2 query user from GROM
	// s3 check user is registered or not
	// s4 check user password is existed or not
	// s5 check user's token is valid or not
	// s6 response about user login
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
}

func GetUserInfo(c *gin.Context) {

}
