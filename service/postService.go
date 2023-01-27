package service

import (
	"errors"
	"web-study/dao/mysql"
	"web-study/entity"
	"web-study/pkg/snowflake"
)

var (
	ErrorPostData      = errors.New("帖子数据不存在")
	ErrorCommunityData = errors.New("社区数据不存在")
)

func CreatPost(data *entity.ParamPostData) (err error) {
	// 生成post id
	data.PostID = snowflake.GenID()
	// 保存数据库
	if err = mysql.InsertPostData(data); err != nil {
		return err
	}
	// 返回
	return nil
}

func GetPostList() (lists []entity.ParamPostData, err error) {
	lists, err = mysql.GetPostList()
	if err != nil {
		return nil, err
	}
	return lists, err
}

func GetPostDataById(id int) (apiData *entity.ApiPostData, err error) {
	data, err := mysql.GetPostData(id)
	if data.PostID == 0 {
		return apiData, ErrorPostData
	}
	if err != nil {
		return apiData, ErrorPostData
	}
	//根据用户id查用户信息
	user := mysql.SelectUserById(int(data.AuthorID))
	//根据社区id查社区信息
	community, err := mysql.FindCommunityById(data.CommunityID)
	if err != nil {
		return apiData, ErrorCommunityData
	}

	apiData = &entity.ApiPostData{
		AuthorName:     user.UserName,
		ParamPostData:  data,
		ParamCommunity: community,
	}

	return apiData, err
}
