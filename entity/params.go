package entity

const (
	OrderTime  = "time"
	OrderScore = "score"
)

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

type ParamCommunity struct {
	CommunityID   int    `json:"communityID" binding:"required"`
	CommunityName string `json:"communityName"  binding:"required"`
	Introduction  string `json:"introduction"  binding:"required"`
}

type ParamListCommunity struct {
	CommunityID   int    `json:"communityID" gorm:"community_id"`
	CommunityName string `json:"communityName" gorm:"community_name"`
}

type ParamPostData struct {
	PostID      int64  `json:"postID,string" `                  //帖子id
	AuthorID    int64  `json:"authorID"`                        //作者id
	CommunityID int    `json:"communityID"  binding:"required"` //社区id
	Title       string `json:"title"  binding:"required"`       //标题
	Content     string `json:"content"  binding:"required"`     //内容
}
type ParamPostDataPlus struct {
	Page  int64  `form:"page"`
	Size  int64  `form:"size"`
	Order string `form:"order" binding:"required"`
}

type ParamVoteData struct {
	// UserID  从请求中获取当前用户
	PostID    string `json:"postID" binding:"required"`               // 帖子id
	Direction int    `json:"direction,string" binding:"oneof=1 0 -1"` // 赞成票（1） 反对票（-1） 取消投票（0）
}
