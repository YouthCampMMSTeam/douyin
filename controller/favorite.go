package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"simple-demo-main/service"
	"strconv"
)

type likeResponse struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

type GetFavouriteListResponse struct {
	StatusCode int32           `json:"status_code"`
	StatusMsg  string          `json:"status_msg,omitempty"`
	VideoList  []service.Video `json:"video_list,omitempty"`
}

// FavoriteAction no practical effect, just check if token is valid
func FavoriteAction(c *gin.Context) {
	token := c.Query("token")

	if _, exist := usersLoginInfo[token]; exist {
		c.JSON(http.StatusOK, Response{StatusCode: 0})
	} else {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}

	strUserId := c.GetString("userId")
	userId, _ := strconv.ParseInt(strUserId, 10, 64)
	strVideoId := c.Query("video_id")
	videoId, _ := strconv.ParseInt(strVideoId, 10, 64)
	strActionType := c.Query("action_type")
	actionType, _ := strconv.ParseInt(strActionType, 10, 64)

	like := new(service.LikeServiceImpl)
	//获取点赞或者取消赞操作的错误信息
	err := like.FavouriteAction(userId, videoId, int32(actionType))
	if err == nil {
		log.Printf("方法like.FavouriteAction(userid, videoId, int32(actiontype) 成功")
		c.JSON(http.StatusOK, likeResponse{
			StatusCode: 0,
			StatusMsg:  "favourite action success",
		})
	} else {
		log.Printf("方法like.FavouriteAction(userid, videoId, int32(actiontype) 失败：%v", err)
		c.JSON(http.StatusOK, likeResponse{
			StatusCode: 1,
			StatusMsg:  "favourite action fail",
		})
	}
}

// FavoriteList all users have same favorite video list
func FavoriteList(c *gin.Context) {
	c.JSON(http.StatusOK, VideoListResponse{
		Response: Response{
			StatusCode: 0,
		},
		VideoList: DemoVideos,
	})
}
