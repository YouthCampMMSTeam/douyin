package model

/*
	喜欢数据模型
*/

type Favorite struct {
	Common
	VideoId   uint64 `json:"video_id"`
	UserId    uint64 `json:"user_id"`
	IsDeleted bool   `json:"is_deleted"`
}
