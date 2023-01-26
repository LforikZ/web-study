package entity

// ParamSignUp 定义请求的参数结构体
type ParamSignUp struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"repassword" binding:"required,eqfield=Password"`
}

type ParamLoginUp struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type ParamInsertCommunity struct {
	CommunityID   int    `json:"communityID" binding:"required"`
	CommunityName string `json:"communityName"  binding:"required"`
	Introduction  string `json:"introduction"  binding:"required"`
}

type ParamListCommunity struct {
	CommunityID   int    `json:"communityID" gorm:"community_id"`
	CommunityName string `json:"communityName" gorm:"community_name"`
}

type ParamPostData struct {
	PostID      int64  `json:"postID" `                         //帖子id
	AuthorID    int64  `json:"authorID"`                        //作者id
	CommunityID int    `json:"communityID"  binding:"required"` //社区id
	Title       string `json:"title"  binding:"required"`       //标题
	Content     string `json:"content"  binding:"required"`     //内容
}
