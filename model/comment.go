package model

/*
	评论存储模型
*/

type Comment struct {
	Common
	VideoId uint64 `json:"video_id"`
	UserId  uint64 `json:"user_id"`
	Content string `json:"content"`
}
