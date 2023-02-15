package model

/*
	消息数据模型
*/

type Message struct {
	Common
	FromId  uint64 `json:"from_id"`
	ToId    uint64 `json:"to_id"`
	Content string `json:"content"`
}
