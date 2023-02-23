package Test

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
var usersLoginInfo = map[string]response.User{"fanyongqi": {
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
		newUser := response.User{
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
