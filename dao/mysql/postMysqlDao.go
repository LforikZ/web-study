package mysql

import (
	"database/sql"
	"errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"strconv"
	"web-study/entity"
)

var ErrIDsNull = errors.New("ids为空值")

type Post struct {
	gorm.Model
	PostID      int64  `json:"postID,string" gorm:"not null;uniqueIndex"`    //帖子id
	AuthorID    int64  `json:"authorID" gorm:"not null"`                     //作者id
	CommunityID int    `json:"communityID" gorm:"not null"`                  //社区id
	Title       string `json:"title" gorm:"size:200;default:'';not null"`    //标题
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

func GetPostList(page, size int) (a []*entity.ParamPostData, err error) {
	var post []Post
	if result := db.Scopes(Paginate(page, size)).Order("updated_at DESC").Find(&post); result.Error == sql.ErrNoRows {
		zap.L().Warn("this is no post in db")
		err = result.Error
	}
	for _, list := range post {
		middleList := &entity.ParamPostData{
			PostID:      list.PostID,
			AuthorID:    list.AuthorID,
			CommunityID: list.CommunityID,
			Title:       list.Title,
			Content:     list.Content,
		}
		a = append(a, middleList)
	}
	return a, err
}

func GetPostListByIDs(ids []string) (a []*entity.ParamPostData, err error) {
	var post []Post
	if len(ids) == 0 {
		return a, ErrIDsNull
	}
	db.Where("post_id IN (?)", ids).Find(&post)
	postMap := make(map[string]Post)
	for _, p := range post {
		postMap[strconv.FormatInt(p.PostID, 10)] = p
	}
	for _, id := range ids {
		if p, ok := postMap[id]; ok {
			middleList := &entity.ParamPostData{
				PostID:      p.PostID,
				AuthorID:    p.AuthorID,
				CommunityID: p.CommunityID,
				Title:       p.Title,
				Content:     p.Content,
			}
			a = append(a, middleList)
		}
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
	data = &entity.ParamPostData{
		PostID:      post.PostID,
		AuthorID:    post.AuthorID,
		CommunityID: post.CommunityID,
		Title:       post.Title,
		Content:     post.Content,
	}

	return data, err
}
