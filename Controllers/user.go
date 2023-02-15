package Controllers

import (
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

/* demo代码，注册、登录、个人信息三个函数，定义了一个map[string]User变量 做测试
// user data will be cleared every time the server starts
// test data: username=zhanglei, password=douyin
func Register(c *gin.Context) {}
func Login(c *gin.Context) {}
func UserInfo(c *gin.Context) {}
*/

// public function
type RegisterRequest struct {
	Username string `json:"username" binding:"required,min=6,max=16" form:"username"`
	Password string `json:"password" binding:"required,min=6,max=16" form:"password"`
}

func hashEncode(str string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(str), bcrypt.DefaultCost)
	if err != nil {
		logrus.Info("加密异常，编码失败:%w", err)
	}
	return string(hash)
}

func comparePasswords(sourcePwd, hashPwd string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashPwd), []byte(sourcePwd)) == nil
}
