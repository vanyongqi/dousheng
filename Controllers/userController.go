package Controllers

import (
	"dousheng-backend/Controllers/common"
	"dousheng-backend/Databases"
	"dousheng-backend/Databases/DAO"
	"dousheng-backend/Middlewares"
	"dousheng-backend/Models"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"net/http"
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

type RegisterLoginRequest struct {
	UserName string `json:"username" binding:"required,min=4,max=16" form:"username"`
	Password string `json:"password" binding:"required,min=4,max=16" form:"password"`
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
	db := Databases.DatabaseSession()
	var req RegisterLoginRequest

	// s1 verify parameters is valid
	if err := c.ShouldBindQuery(&req); err != nil {
		logrus.Error("用户信息绑定失败,Tips：ID长度为4-16个字，请核对！")
		c.JSON(http.StatusUnprocessableEntity, common.Response{
			StatusCode: 1,
			StatusMsg:  "用户信息绑定失败，Tips：ID长度为4-16个字，请核对！",
		})
		c.JSON(200, gin.H{"name": req.UserName, "password": req.Password})
		return
	}

	// s2 verify user is already registered or not
	if DAO.GetUserByName(db, req.UserName) != nil {
		logrus.Warn("用户名重复，用户注册失败！the name already exist in database")
		c.JSON(http.StatusOK, common.Response{
			StatusCode: 2,
			StatusMsg:  "用户名重复，用户注册失败！the name already exist in database",
		})
		c.JSON(200, req.UserName)
		return
	}
	// s3 create user from GROM by username and password
	registerUser := DAO.CreateUser(db, &Models.User{
		Name:     req.UserName,
		Password: BcryptEncode(req.Password),
		Content:  "",
	})
	// s4 generate token from middleware layer
	token, err := Middlewares.CreateToken(registerUser.ID, registerUser.Name)
	if err != nil {
		logrus.Info(" error happened while creating token")
		c.JSON(http.StatusInternalServerError, common.Response{
			StatusCode: 3,
			StatusMsg:  "error happened while creating token",
		})
	}
	// s5 response about user's registration
	c.JSON(http.StatusOK, common.UserLoginRegisterResponse{
		Response: common.Response{
			StatusCode: 0,
			StatusMsg:  "success",
		},
		UserID: int64(registerUser.ID),
		Token:  token,
	})

}

func PostLoginUser(c *gin.Context) {
	db := Databases.DatabaseSession()
	var req RegisterLoginRequest

	// s1 verify parameters is valid
	if err := c.ShouldBindQuery(&req); err != nil {
		logrus.Error(" 登录中用户信息绑定失败")
		c.JSON(http.StatusUnprocessableEntity, common.Response{
			StatusCode: 1,
			StatusMsg:  " 登录中用户信息绑定失败",
		})
		c.JSON(200, gin.H{"name": req.UserName, "password": req.Password})
		return
	}
	// s2 query user from GROM
	loginUser := DAO.GetUserByName(db, req.UserName)

	// s3 check user is registered or not
	if loginUser == nil {
		c.JSON(http.StatusExpectationFailed, common.Response{
			StatusCode: http.StatusExpectationFailed,
			StatusMsg:  "用户不存在",
		})
		return
	}

	// s4 check user password is existed or not
	if VerifyPassword(req.Password, loginUser.Password) == false {
		c.JSON(http.StatusExpectationFailed, common.Response{
			StatusCode: http.StatusExpectationFailed,
			StatusMsg:  "用户密码错误",
		})
		c.JSON(http.StatusExpectationFailed, gin.H{
			"name":          req.UserName,
			"password":      req.Password,
			"loginPassword": loginUser.Password,
		})
		return
	}
	// s5 check user's token is valid or not
	token, err := Middlewares.CreateToken(loginUser.ID, loginUser.Name)
	if err != nil {
		c.JSON(http.StatusExpectationFailed, common.Response{
			StatusCode: http.StatusExpectationFailed,
			StatusMsg:  "token创建失败，error happened while creating token",
		})
		return
	}
	// s6 response about user login
	c.JSON(http.StatusOK, common.UserLoginRegisterResponse{
		Response: common.Response{
			StatusCode: 0,
			StatusMsg:  "success",
		},
		UserID: int64(loginUser.ID),
		Token:  token,
	})
}

func GetUserInfo(c *gin.Context) {

}
