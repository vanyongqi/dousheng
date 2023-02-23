package Utils

import (
	"dousheng-backend/Controllers/response"
	"dousheng-backend/Models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func QueryIDFromCtx(c *gin.Context, queryName string) uint {
	id, err := strconv.ParseUint(c.Query(queryName), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 1,
			StatusMsg:  "解析id异常",
		})
		return 0
	}
	return uint(id)
}

func GetUserFromCtx(c *gin.Context) *Models.User {
	user_, _ := c.Get("User")
	//fmt.Print(" func utils problem")
	//fmt.Println(user_)
	user, _ := user_.(*Models.User)
	return user
}
