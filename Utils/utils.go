package Utils

import (
	"dousheng-backend/Controllers/common"
	"dousheng-backend/Models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func BindAndValid(c *gin.Context, target interface{}) bool {
	var queryOK bool = true

	if err := c.ShouldBindQuery(target); err != nil {
		queryOK = false
	}

	return queryOK
}

func QueryIDAndValid(ctx *gin.Context, queryName string) uint {
	id, err := strconv.ParseUint(ctx.Query(queryName), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, common.Response{
			StatusCode: 1,
			StatusMsg:  queryName + "不是数字",
		})
		return 0
	}
	return uint(id)
}

func GetUserFromCtx(ctx *gin.Context) *Models.User {
	user_, _ := ctx.Get("User")
	user, _ := user_.(*Models.User)
	return user
}
