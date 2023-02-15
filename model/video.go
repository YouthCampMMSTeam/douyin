package model

/*
	视频存储模型
*/

import (
	"time"
)

type Video struct {
	Common
	AuthorId      uint64 `json:"author_id"`
	PlayUrl       string `json:"play_url"`
	CoverUrl      string `json:"cover_url"`
	FavoriteCount uint64 `json:"favorite_count"`
	CommentCount  uint64 `json:"comment_count"`
	Title         string `json:"title"`
}

//查询视频列表(根据视频发布时间，查询最新的30条视频)
func GetVideoListByCreateTime(latest_time int64) ([]*Video, error) {
	var videoList []*Video
	latest := time.UnixMilli(latest_time)
	err := DB.Limit(30).Where("create_at < ?", latest).Find(&videoList).Error
	if err != nil {
		return nil, err
	}
	return videoList, nil
}
