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
	CommunityID   int    `json:"communityID" gorm:"community_id" binding:"required"`
	CommunityName string `json:"communityName" gorm:"community_name" binding:"required"`
	Introduction  string `json:"introduction" gorm:"introduction" binding:"required"`
}

type ParamListCommunity struct {
	CommunityID   int    `json:"communityID" gorm:"community_id"`
	CommunityName string `json:"communityName" gorm:"community_name"`
}
