package Controllers

import (
	"dousheng-backend/Controllers/response"
	"dousheng-backend/Databases"
	"dousheng-backend/Databases/DAO"
	"dousheng-backend/Middlewares"
	"dousheng-backend/Models"
	"dousheng-backend/Models/model2response"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os/exec"
	"path/filepath"
	"time"
	"unicode/utf8"
)

//POST /douyin/publish/action/
//登录用户选择视频上传
//data	file	是	视频数据
//token	string	是	用户鉴权token
//title	string	是	视频标题

func PostPublishAction(c *gin.Context) {
	db := Databases.DatabaseSession()

	token := c.PostForm("token")
	userID, err := Middlewares.ParseToken(c, token)
	if err != nil {
		return
	}

	data, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: response.BADREQUEST,
			StatusMsg:  "文件获取错误: " + err.Error(),
		})
		return
	}

	if data.Filename[len(data.Filename)-3:] != "mp4" {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: response.BADREQUEST,
			StatusMsg:  "不支持的文件格式",
		})
		return
	}

	title := c.PostForm("title")
	if title == "" || utf8.RuneCountInString(title) > 20 {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: response.BADREQUEST,
			StatusMsg:  "标题获取错误",
		})
		return
	}

	var time_stamp int64 = time.Now().Unix()
	fileName := fmt.Sprintf("%d_%s", time_stamp, title+".mp4")
	saveFile := filepath.Join("./public/video/", fileName)

	if err := c.SaveUploadedFile(data, saveFile); err != nil {
		c.JSON(http.StatusInternalServerError, response.Response{
			StatusCode: response.INTERNALERROR,
			StatusMsg:  err.Error(),
		})
		return
	}

	cmd := exec.Command("ffmpeg", "-i", "public/video/"+fileName,
		"-frames:v", "1", "-f", "image2",
		"public/cover4video/"+fileName[:len(fileName)-4]+".jpg")
	if err := cmd.Run(); err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
		c.JSON(http.StatusInternalServerError, response.Response{
			StatusCode: response.INTERNALERROR,
			StatusMsg:  err.Error(),
		})
		return
	}

	DAO.CreateVideo(db, &Models.Video{
		AuthorID: userID,
		Title:    fileName,
	})

	c.JSON(http.StatusOK, response.Response{
		StatusCode: response.SUCCESS,
		StatusMsg:  fileName + " 上传成功",
	})
}

//GET	/douyin/publish/list/
//用户的视频发布列表，直接列出用户所有投稿过的视频
//token	string	是	用户鉴权token
//user_id	query	string	是	用户id

func GetPublishList(c *gin.Context) {
	db := Databases.DatabaseSession()

	//userReq, _ := c.Get("User")
	//user, _ := userReq.(*Models.User)
	//	if c == nil {
	//	fmt.Printf("!!!!!!!!! ctx wrong" + user.Name)
	//	}
	//if user == nil {
	//		fmt.Printf("!!!!!!!!! user wrong")
	//	}

	token := c.Query("token")
	userID, err := Middlewares.ParseToken(c, token)
	if err != nil {
		return
	}
	videos := DAO.GetUserPublishVideosByID(db, userID)
	//fmt.Print(videos)
	//fmt.Print("++++++++++")
	responseVideos := model2response.VideoModeltoRes(db, videos)
	for i := 0; i < len(responseVideos); i++ {
		responseVideos[i].IsFavorite = DAO.IsUserFavoriteVideo(db, userID, uint(responseVideos[i].ID))
	}
	c.JSON(http.StatusOK, response.VideoListResponse{
		Response: response.Response{
			StatusCode: response.SUCCESS,
			StatusMsg:  "获取成功",
		},
		VideoList: responseVideos,
	})
}
