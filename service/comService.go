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

func InsertComData(community *entity.ParamInsertCommunity) (err error) {
	if result, _ := mysql.FindCommunityById(community.CommunityID); result.ID != 0 {
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

func GetCommunityById(id int) (community entity.ParamInsertCommunity, err error) {
	result, err := mysql.FindCommunityById(id)
	if err != nil {
		return community, err
	}
	community.CommunityID = result.CommunityID
	community.CommunityName = result.CommunityName
	community.Introduction = result.Introduction

	return community, err
}
