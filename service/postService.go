package service

import (
	"errors"
	"fmt"
	"web-study/dao/mysql"
	"web-study/dao/redis"
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
	// 保存redis
	if err := redis.CreatPost(data.PostID); err != nil {
		fmt.Println(err)
		return err
	}
	// 保存至mysql
	if err = mysql.InsertPostData(data); err != nil {
		return err
	}

	// 返回
	return nil
}

func GetPostList(page, size int) (lists []*entity.ApiPostData, err error) {
	data, err := mysql.GetPostList(page, size)
	if err != nil {
		return nil, err
	}
	for _, datum := range data {
		//根据用户id查用户信息
		user := mysql.SelectUserById(int(datum.AuthorID))
		//根据社区id查社区信息
		community, err := mysql.FindCommunityById(datum.CommunityID)
		if err != nil {
			return lists, ErrorCommunityData
		}

		middle := &entity.ApiPostData{
			AuthorName:     user.UserName,
			ParamPostData:  datum,
			ParamCommunity: community,
		}
		lists = append(lists, middle)
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
