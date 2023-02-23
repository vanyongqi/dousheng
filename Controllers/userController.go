package Controllers

import (
	"dousheng-backend/Controllers/request"
	"dousheng-backend/Controllers/response"
	"dousheng-backend/Databases"
	"dousheng-backend/Databases/DAO"
	"dousheng-backend/Middlewares"
	"dousheng-backend/Models"
	"dousheng-backend/Models/model2response"
	"dousheng-backend/Utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

// POST /douyin/user/register/
// 参数名	  位置	   类型		必填		说明
// username  query   string  	 是 	注册用户名，最长32个字符
// password  query   string	    是 		密码，最长32个字符
//type RegisterLoginRequest struct {
//	UserName string `json:"username" binding:"required,min=4,max=32" form:"username"`
//	Password string `json:"password" binding:"required,min=4,max=32" form:"password"`
//}

func PostRegisterUser(c *gin.Context) {
	db := Databases.DatabaseSession()
	var req request.RegisterLoginRequest

	// s1 verify parameters is valid
	if err := c.ShouldBindQuery(&req); err != nil {
		logrus.Error("用户信息绑定失败,Tips：ID长度为4-16个字，请核对！")
		c.JSON(http.StatusUnprocessableEntity, response.Response{
			StatusCode: 1,
			StatusMsg:  "用户信息绑定失败，Tips：ID长度为4-16个字，请核对！",
		})
		c.JSON(200, gin.H{"name": req.UserName, "password": req.Password})
		return
	}

	// s2 verify user is already registered or not
	if DAO.GetUserByName(db, req.UserName) != nil {
		logrus.Warn("用户名重复，用户注册失败！the name already exist in database")
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 2,
			StatusMsg:  "用户名重复，用户注册失败！the name already exist in database",
		})
		c.JSON(2, req.UserName)
		return
	}
	// s3 create user from GROM by username and password
	registerUser := DAO.CreateUser(db, &Models.User{
		Name:            req.UserName,
		Password:        Utils.BcryptEncode(req.Password),
		Signature:       "签名写不出对你的心",
		Avatar:          "/default_avatar.jpg",
		BackgroundImage: "/default_background.jpg",
	})
	// s4 generate token from middleware layer
	token, err := Middlewares.CreateToken(registerUser.ID)
	if err != nil {
		logrus.Info(" error happened while creating token")
		c.JSON(http.StatusInternalServerError, response.Response{
			StatusCode: 3,
			StatusMsg:  "error happened while creating token",
		})
	}
	// s5 response about user's registration
	c.JSON(http.StatusOK, response.UserLoginRegisterResponse{
		Response: response.Response{
			StatusCode: 0,
			StatusMsg:  "success",
		},
		UserID: int64(registerUser.ID),
		Token:  token,
	})

}

func PostLoginUser(c *gin.Context) {
	db := Databases.DatabaseSession()
	var req request.RegisterLoginRequest

	// s1 verify parameters is valid
	if err := c.ShouldBindQuery(&req); err != nil {
		logrus.Error(" 登录中用户信息绑定失败")
		c.JSON(http.StatusUnprocessableEntity, response.Response{
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
		c.JSON(http.StatusExpectationFailed, response.Response{
			StatusCode: http.StatusExpectationFailed,
			StatusMsg:  "用户不存在",
		})
		return
	}

	// s4 check user password is existed or not
	if Utils.VerifyPassword(req.Password, loginUser.Password) == false {
		c.JSON(http.StatusExpectationFailed, response.Response{
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
	token, err := Middlewares.CreateToken(loginUser.ID)
	if err != nil {
		c.JSON(http.StatusExpectationFailed, response.Response{
			StatusCode: http.StatusExpectationFailed,
			StatusMsg:  "token创建失败，error happened while creating token",
		})
		return
	}
	// s6 response about user login
	c.JSON(http.StatusOK, response.UserLoginRegisterResponse{
		Response: response.Response{
			StatusCode: 0,
			StatusMsg:  "success",
		},
		UserID: int64(loginUser.ID),
		Token:  token,
	})
}

//GET/douyin/user/
//获取用户的 id、昵称，如果实现社交部分的功能，还会返回关注数和粉丝数

func GetUserInfo(c *gin.Context) {
	db := Databases.DatabaseSession()

	var nowUser *Models.User

	token := c.Query("token")
	if token != "" {
		userID, err := Middlewares.ParseToken(c, token)

		if err != nil {
			return
		}
		nowUser, err = DAO.GetUserByID(db, userID)
		if err != nil {
			c.JSON(http.StatusNotFound, response.Response{
				StatusCode: 3,
				StatusMsg:  "找不到用户",
			})
			return
		}
	}

	toUserID := Utils.QueryIDFromCtx(c, "user_id")
	if toUserID == 0 {
		return
	}

	toUser, err := DAO.GetUserByID(db, toUserID)
	if err != nil {
		c.JSON(http.StatusNotFound, response.Response{
			StatusCode: 3,
			StatusMsg:  "找不到用户",
		})
		return
	}
	responseUser := model2response.UserModeltoRe(db, *toUser)
	if nowUser != nil {
		responseUser.IsFollow = DAO.IsUserFollow(db, nowUser.ID, toUserID)
	} else {
		responseUser.IsFollow = false
	}

	c.JSON(http.StatusOK, response.UserResponse{
		Response: response.Response{
			StatusCode: 0,
			StatusMsg:  "已找到用户",
		},
		User: responseUser,
	})

}
