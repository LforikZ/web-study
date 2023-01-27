package entity

import "web-study/dao/mysql"

// ApiPostData 帖子详情接口结构体
type ApiPostData struct {
	AuthorName       string `json:"authorName"`
	*mysql.Post             //嵌入帖子结构体
	*mysql.Community        //嵌入社区结构体
}
