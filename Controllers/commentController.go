package Controllers

import (
	"dousheng-backend/Controllers/request"
	"dousheng-backend/Controllers/response"
	"dousheng-backend/Databases"
	"dousheng-backend/Databases/DAO"
	"dousheng-backend/Models"
	"dousheng-backend/Models/model2response"
	"dousheng-backend/Utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func PostCommentAction(c *gin.Context) {
	db := Databases.DatabaseSession()
	var req request.CommentRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		logrus.Warn("var wrong binding")
		return
	}

	user := Utils.GetUserFromCtx(c)
	if req.ActionType == 1 {
		if len(req.CommentText) == 0 {
			c.JSON(http.StatusBadRequest, response.Response{
				StatusCode: response.BADREQUEST,
				StatusMsg:  "无内容",
			})
			return
		}
		comment := DAO.CreateComment(db, &Models.Comment{
			UserID:  user.ID,
			VideoID: req.VideoID,
			Content: req.CommentText,
		})
		c.JSON(http.StatusOK, response.CommentResponse{
			Response: response.Response{
				StatusCode: response.SUCCESS,
				StatusMsg:  "评论成功",
			},
			Comment: model2response.CommentModeltoRe(db, *comment),
		})
	} else {
		if req.CommentID == 0 {
			c.JSON(http.StatusBadRequest, response.Response{
				StatusCode: response.BADREQUEST,
				StatusMsg:  "评论异常",
			})
			return
		}
		if err := DAO.DeleteComment(db, req.CommentID); err != nil {
			db.Rollback()
			c.JSON(http.StatusNotFound, response.Response{
				StatusCode: response.NOTFOUND,
				StatusMsg:  "评论异常",
			})
			return
		}
		c.JSON(http.StatusOK, response.Response{
			StatusCode: response.SUCCESS,
			StatusMsg:  "删除成功",
		})
	}

}
func GetCommentList(c *gin.Context) {
	db := Databases.DatabaseSession()

	videoID := Utils.QueryIDFromCtx(c, "video_id")
	if videoID == 0 {
		return
	}

	c.JSON(http.StatusOK, response.CommentListResponse{
		Response: response.Response{
			StatusCode: response.SUCCESS,
			StatusMsg:  "评论列表获取成功",
		},
		CommentList: model2response.CommentModeltoRes(db, DAO.GetComments(db, videoID)),
	})
}
