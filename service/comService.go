package service

import (
	"errors"
	"web-study/dao/mysql"
	"web-study/entity"
)

var (
	ErrorComIDExit   = errors.New("社区ID已存在")
	ErrorComNameExit = errors.New("社区名字已存在")
)

func InsertComData(community *entity.ParamCommunity) (err error) {
	if result, _ := mysql.FindCommunityById(community.CommunityID); result.CommunityID != 0 {
		return ErrorComIDExit
	}
	if result := mysql.FindCommunityByName(community.CommunityName); result.ID != 0 {
		return ErrorComNameExit
	}

	if err = mysql.InsertCommunityData(community); err != nil {
		return err
	}
	return
}

func GetCommunityList() ([]entity.ParamListCommunity, error) {
	list, err := mysql.FindCommunityList()
	return list, err
}

func GetCommunityById(id int) (community *entity.ParamCommunity, err error) {
	result, err := mysql.FindCommunityById(id)
	if err != nil {
		return community, err
	}

	community = &entity.ParamCommunity{
		CommunityID:   result.CommunityID,
		CommunityName: result.CommunityName,
		Introduction:  result.Introduction,
	}

	return community, err
}
