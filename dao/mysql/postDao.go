package mysql

import (
	"gorm.io/gorm"
	"web-study/entity"
)

type Post struct {
	gorm.Model
	PostID      int64  `json:"postID" gorm:"not null;uniqueIndex"`           //帖子id
	AuthorID    int64  `json:"authorID" gorm:"not null"`                     //作者id
	CommunityID int    `json:"communityID" gorm:"not null"`                  //社区id
	Title       string `json:"title" gorm:"size:20;default:'';not null"`     //标题
	Content     string `json:"content" gorm:"size:8888;not null;default:''"` //内容
}

func InsertPostData(data *entity.ParamPostData) error {
	post := Post{
		PostID:      data.PostID,
		AuthorID:    data.AuthorID,
		CommunityID: data.CommunityID,
		Title:       data.Title,
		Content:     data.Content,
	}
	if result := db.Create(&post); result.Error != nil {
		return result.Error
	}
	return nil
}
