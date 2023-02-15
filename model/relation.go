package model

/*
	关系数据模型
*/

type Relation struct {
	Common
	FollowId   int64 `json:"follow_id"`
	FollowerId int64 `json:"follower_id"`
}
