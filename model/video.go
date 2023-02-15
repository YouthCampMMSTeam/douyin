package model

/*
	视频存储模型
*/

type Video struct {
	Common
	AuthorId      uint64 `json:"author_id"`
	PlayUrl       string `json:"play_url"`
	CoverUrl      string `json:"cover_url"`
	FavoriteCount uint64 `json:"favorite_count"`
	CommentCount  uint64 `json:"comment_count"`
	Title         string `json:"title"`
}
