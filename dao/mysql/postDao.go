package mysql

import (
	"database/sql"
	"go.uber.org/zap"
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

func GetPostList() (a []entity.ParamPostData, err error) {
	var post []Post
	if result := db.Find(&post); result.Error == sql.ErrNoRows {
		zap.L().Warn("this is no post in db")
		err = result.Error
	}
	for _, list := range post {
		var middleList entity.ParamPostData
		middleList.PostID = list.PostID
		middleList.Title = list.Title
		middleList.CommunityID = list.CommunityID
		middleList.Content = list.Content
		middleList.AuthorID = list.AuthorID
		a = append(a, middleList)
	}
	return a, err
}

func GetPostData(id int) (data *entity.ParamPostData, err error) {
	var post Post
	if result := db.Where("post_id=?", id).Find(&post); result.Error == sql.ErrNoRows {
		zap.L().Warn("this is no post in db")
		err = result.Error
		return data, err
	}

	data.PostID = post.PostID
	data.Title = post.Title
	data.CommunityID = post.CommunityID
	data.Content = post.Content
	data.AuthorID = post.AuthorID

	return data, err
}
