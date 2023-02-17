package Controllers

import (
	"dousheng-backend/Controllers/common"
	"dousheng-backend/Databases"
	"dousheng-backend/Databases/DAO"
	"dousheng-backend/Middlewares"
	"dousheng-backend/Models"
	"dousheng-backend/Utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

func PostRegisterUser(c *gin.Context) {
	db := Databases.DatabaseSession()
	var req common.RegisterLoginRequest

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
		c.JSON(http.StatusBadRequest, common.Response{
			StatusCode: 2,
			StatusMsg:  "用户名重复，用户注册失败！the name already exist in database",
		})
		c.JSON(2, req.UserName)
		return
	}
	// s3 create user from GROM by username and password
	registerUser := DAO.CreateUser(db, &Models.User{
		Name:     req.UserName,
		Password: Utils.BcryptEncode(req.Password),
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
	var req common.RegisterLoginRequest

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
	if Utils.VerifyPassword(req.Password, loginUser.Password) == false {
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
	db := Databases.DatabaseSession()
	var loginUser *Models.User
	token := c.Query("token")
	if token != "" {
		loginUserID, loginUserName, err := Middlewares.ParseToken(c, token)
		if err != nil {
			return
		}
		loginUser, err = DAO.GetUserByID(db, loginUserID)
		logrus.Info(string(loginUserName) + " 用户尝试登录 in" + time.Now().String())
		if err != nil {
			c.JSON(http.StatusExpectationFailed, common.Response{
				StatusCode: http.StatusExpectationFailed,
				StatusMsg:  "用户不存在",
			})
			return
		}

	}

	c.JSON(http.StatusOK, loginUser)

}
