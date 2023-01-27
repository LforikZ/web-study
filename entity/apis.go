package entity

// ApiPostData 帖子详情接口结构体
type ApiPostData struct {
	AuthorName      string             `json:"authorName"`
	*ParamPostData                     //嵌入帖子结构体
	*ParamCommunity `json:"community"` //嵌入社区结构体
}
