package entity

// ApiPostData 帖子详情接口结构体
type ApiPostData struct {
	AuthorName      string             `json:"authorName"`
	*ParamPostData                     //嵌入帖子结构体
	*ParamCommunity `json:"community"` //嵌入社区结构体
}

type ApiUserData struct {
	UserID   int64  `json:"userID,string"` //用户id
	UserName string `json:"userName"`      //用户名字
	Token    string `json:"token"`         //token
}
